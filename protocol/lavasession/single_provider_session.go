package lavasession

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lavanet/lava/utils"
)

const (
	TRY_LOCK_ATTEMPTS = 30
)

type SingleProviderSession struct {
	userSessionsParent *ProviderSessionsWithConsumer
	CuSum              uint64
	LatestRelayCu      uint64
	SessionID          uint64
	lock               sync.RWMutex
	RelayNum           uint64
	PairingEpoch       uint64
	occupyingGuid      uint64 //used for tracking errors
}

// to be used only when locked, otherwise can return wrong values
// is used to determine if the proof is beneficial and needs to be sent to rewardServer
func (sps *SingleProviderSession) IsPayingRelay() bool {
	return sps.LatestRelayCu > 0
}

func (sps *SingleProviderSession) writeCuSumAtomically(cuSum uint64) {
	atomic.StoreUint64(&sps.CuSum, cuSum)
}

func (sps *SingleProviderSession) atomicReadCuSum() uint64 {
	return atomic.LoadUint64(&sps.CuSum)
}

func (sps *SingleProviderSession) lockForUse(ctx context.Context) {
	guid, found := utils.GetUniqueIdentifier(ctx)
	sps.lock.Lock()
	if found {
		sps.setOccupyingGuid(guid)
	}
}

func (sps *SingleProviderSession) tryLockForUse(ctx context.Context) error {
	guid, found := utils.GetUniqueIdentifier(ctx)
	locked := false
	for i := 0; i < TRY_LOCK_ATTEMPTS; i++ {
		locked = sps.lock.TryLock()
		if locked {
			break
		}
		time.Sleep(time.Millisecond)
	}
	if locked && found {
		sps.setOccupyingGuid(guid)
	}
	if locked {
		return nil
	}
	occupyingGuid := sps.GetOccupyingGuid()
	return utils.LavaFormatError("tryLockForUse failure", LockMisUseDetectedError, utils.Attribute{Key: "GUID", Value: ctx}, utils.Attribute{Key: "occupyingGuid", Value: occupyingGuid})
}

func (sps *SingleProviderSession) GetOccupyingGuid() uint64 {
	return atomic.LoadUint64(&sps.occupyingGuid)
}

func (sps *SingleProviderSession) setOccupyingGuid(occupyingGuid uint64) {
	atomic.StoreUint64(&sps.occupyingGuid, occupyingGuid)
}

func (sps *SingleProviderSession) GetPairingEpoch() uint64 {
	return atomic.LoadUint64(&sps.PairingEpoch)
}

func (sps *SingleProviderSession) SetPairingEpoch(epoch uint64) {
	atomic.StoreUint64(&sps.PairingEpoch, epoch)
}

// Verify the SingleProviderSession is locked when getting to this function, if its not locked throw an error
func (sps *SingleProviderSession) VerifyLock() error {
	if sps.lock.TryLock() { // verify.
		// if we managed to lock throw an error for misuse.
		defer sps.lock.Unlock()
		return utils.LavaFormatError("verifyLock failure, lock was free", LockMisUseDetectedError)
	}
	return nil
}

// In case the user session is a data reliability we just need to verify that the cusum is the amount agreed between the consumer and the provider
func (sps *SingleProviderSession) PrepareDataReliabilitySessionForUsage(relayRequestTotalCU uint64) error {
	if relayRequestTotalCU != DataReliabilityCuSum {
		return utils.LavaFormatError("PrepareDataReliabilitySessionForUsage", DataReliabilityCuSumMisMatchError, utils.Attribute{Key: "relayRequestTotalCU", Value: relayRequestTotalCU})
	}
	sps.LatestRelayCu = DataReliabilityCuSum // 1. update latest
	sps.CuSum = relayRequestTotalCU          // 2. update CuSum, if consumer wants to pay more, let it
	utils.LavaFormatDebug("PrepareDataReliabilitySessionForUsage",
		utils.Attribute{Key: "relayRequestTotalCU", Value: relayRequestTotalCU},
		utils.Attribute{Key: "sps.LatestRelayCu", Value: sps.LatestRelayCu},
	)
	return nil
}

func (sps *SingleProviderSession) PrepareSessionForUsage(ctx context.Context, cuFromSpec uint64, relayRequestTotalCU uint64, allowedThreshold float64) error {
	err := sps.VerifyLock() // sps is locked
	if err != nil {
		return utils.LavaFormatError("sps.verifyLock() failed in PrepareSessionForUsage", err, utils.Attribute{Key: "GUID", Value: ctx}, utils.Attribute{Key: "relayNum", Value: sps.RelayNum}, utils.Attribute{Key: "sps.sessionId", Value: sps.SessionID})
	}

	// checking if this user session is a data reliability user session.
	if sps.userSessionsParent.atomicReadIsDataReliability() == isDataReliabilityPSWC {
		return sps.PrepareDataReliabilitySessionForUsage(relayRequestTotalCU)
	}

	maxCu := sps.userSessionsParent.atomicReadMaxComputeUnits()
	if relayRequestTotalCU < sps.CuSum+cuFromSpec {
		// there is a mismatch, check if it's critical
		// there are allowed cases when a mismatch happens
		// 1) mismatch still provides us with more CU, count the missing diff as missing, we still send the new proof to reward server
		// 2) mismatch provides us with less total CU, count all of the request CU as missing, and do not send the proof (setting sps.LatestRelayCu to 0)

		missingCU := cuFromSpec
		if relayRequestTotalCU > sps.CuSum {
			// case 1) expected: cuFromSpec + sps.CuSum, given: relayRequestTotalCU, missing: expected-given
			missingCU = cuFromSpec + sps.CuSum - relayRequestTotalCU
		} else {
			// case 2) the relay is giving less than our latest proof, it's missing the entire spec cu
			relayRequestTotalCU = sps.CuSum // sets cuToAdd to 0
		}

		var cuErr error = nil
		// verify there are enough missing cus allowed
		canAddMissingCU := sps.userSessionsParent.SafeAddMissingComputeUnits(missingCU, allowedThreshold)
		if !canAddMissingCU {
			cuErr = utils.LavaFormatWarning("CU mismatch PrepareSessionForUsage, Provider and consumer disagree on CuSum", ProviderConsumerCuMisMatch,
				utils.Attribute{Key: "request.CuSum", Value: relayRequestTotalCU},
				utils.Attribute{Key: "provider.CuSum", Value: sps.CuSum},
				utils.Attribute{Key: "specCU", Value: cuFromSpec},
				utils.Attribute{Key: "expected", Value: sps.CuSum + cuFromSpec},
				utils.Attribute{Key: "GUID", Value: ctx},
				utils.Attribute{Key: "relayNum", Value: sps.RelayNum},
				utils.Attribute{Key: "missingCUs", Value: missingCU},
				utils.Attribute{Key: "allowedThreshold", Value: allowedThreshold},
			)
		}
		// verify missing cus aren't immediately expended and are scattered across the session duration

		if cuErr != nil {
			sps.lock.Unlock() // unlock on error
			return cuErr
		}
		// there are missing CU but that's fine because it's within the threshold, and provider gets paid for the new request
		// reading userSessionParent address because it's a fixed string value that isn't changing
		utils.LavaFormatWarning("CU Mismatch within the threshold", nil, utils.Attribute{Key: "GUID", Value: ctx}, utils.Attribute{Key: "missingCU", Value: missingCU}, utils.Attribute{Key: "consumer", Value: sps.userSessionsParent.consumerAddr},
			utils.Attribute{Key: "sessionID", Value: sps.SessionID}, utils.Attribute{Key: "relayNum", Value: sps.RelayNum})
	}

	// if consumer wants to pay more, we need to adjust the payment. so next relay will be in sync
	cuToAdd := relayRequestTotalCU - sps.CuSum // how much consumer thinks he needs to pay - our current state

	// this must happen first, as we also validate and add the used cu to parent here
	err = sps.validateAndAddUsedCU(cuToAdd, maxCu)
	if err != nil {
		sps.lock.Unlock() // unlock on error
		return err
	}
	// finished validating, can add all info.
	sps.LatestRelayCu = cuToAdd // 1. update latest
	sps.CuSum += cuToAdd        // 2. update CuSum, if consumer wants to pay more, let it
	utils.LavaFormatDebug("Before Update Normal PrepareSessionForUsage",
		utils.Attribute{Key: "GUID", Value: ctx},
		utils.Attribute{Key: "relayRequestTotalCU", Value: relayRequestTotalCU},
		utils.Attribute{Key: "sps.LatestRelayCu", Value: sps.LatestRelayCu},
		utils.Attribute{Key: "sps.CuSum", Value: sps.CuSum},
		utils.Attribute{Key: "sps.sessionId", Value: sps.SessionID},
		utils.Attribute{Key: "relayNum", Value: sps.RelayNum},
	)
	return nil
}

func (sps *SingleProviderSession) validateAndAddUsedCU(currentCU uint64, maxCu uint64) error {
	for {
		usedCu := sps.userSessionsParent.atomicReadUsedComputeUnits() // check used cu now
		if usedCu+currentCU > maxCu {
			return utils.LavaFormatError("Maximum cu exceeded PrepareSessionForUsage", MaximumCULimitReachedByConsumer,
				utils.Attribute{Key: "usedCu", Value: usedCu},
				utils.Attribute{Key: "currentCU", Value: currentCU},
				utils.Attribute{Key: "maxCu", Value: maxCu},
			)
		}
		// compare usedCu + current cu vs usedCu, if swap succeeds, return otherwise try again
		// this can happen when multiple sessions are adding their cu at the same time.
		// comparing and adding is protecting against race conditions as the parent is not locked.
		if sps.userSessionsParent.atomicCompareAndWriteUsedComputeUnits(usedCu+currentCU, usedCu) {
			return nil
		}
	}
}

func (sps *SingleProviderSession) validateAndSubUsedCU(currentCU uint64) error {
	for {
		usedCu := sps.userSessionsParent.atomicReadUsedComputeUnits()                               // check used cu now
		if sps.userSessionsParent.atomicCompareAndWriteUsedComputeUnits(usedCu-currentCU, usedCu) { // decrease the amount of used cu from the known value
			return nil
		}
	}
}

// for a different behavior in data reliability session failure add here
func (sps *SingleProviderSession) onDataReliabilitySessionFailure() error {
	return nil
}

func (sps *SingleProviderSession) onSessionFailure() error {
	err := sps.VerifyLock() // sps is locked
	if err != nil {
		return utils.LavaFormatError("sps.verifyLock() failed in onSessionFailure", err, utils.Attribute{Key: "sessionID", Value: sps.SessionID})
	}
	defer sps.lock.Unlock()

	// handle data reliability session failure
	if sps.userSessionsParent.atomicReadIsDataReliability() == isDataReliabilityPSWC {
		return sps.onDataReliabilitySessionFailure()
	}

	sps.CuSum -= sps.LatestRelayCu
	sps.validateAndSubUsedCU(sps.LatestRelayCu)
	sps.LatestRelayCu = 0
	return nil
}

func (sps *SingleProviderSession) onSessionDone(relayNumber uint64) error {
	// this can be called on collected sessions, so if in the future you need to touch the parent, take this into consideration to change the OnSessionDone calls in provider_session_manager
	err := sps.VerifyLock() // sps is locked
	if err != nil {
		return utils.LavaFormatError("sps.verifyLock() failed in onSessionDone", err)
	}
	sps.RelayNum = relayNumber
	sps.LatestRelayCu = 0 // reset the cu, we can also verify its 0 when loading.
	sps.lock.Unlock()
	return nil
}
