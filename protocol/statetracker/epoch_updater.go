package statetracker

import (
	"sync"
	"sync/atomic"

	"github.com/lavanet/lava/utils"
	"golang.org/x/net/context"
)

const (
	CallbackKeyForEpochUpdate = "epoch-update"
)

type EpochUpdatable interface {
	UpdateEpoch(epoch uint64)
	UpdateVirtualEpoch(epoch uint64, virtualEpoch uint64)
}

type EpochUpdater struct {
	lock                sync.RWMutex
	epochUpdatables     []*EpochUpdatable
	currentEpoch        uint64
	currentVirtualEpoch uint64
	stateQuery          *EpochStateQuery
}

func NewEpochUpdater(stateQuery *EpochStateQuery) *EpochUpdater {
	return &EpochUpdater{epochUpdatables: []*EpochUpdatable{}, stateQuery: stateQuery}
}

func (eu *EpochUpdater) RegisterEpochUpdatable(ctx context.Context, epochUpdatable EpochUpdatable) {
	eu.lock.Lock()
	defer eu.lock.Unlock()
	// initialize with the current epoch
	currentEpoch, err := eu.stateQuery.CurrentEpochStart(ctx)
	if err != nil {
		utils.LavaFormatFatal("epoch updatable failed registering for epoch updates", err)
	}
	epochUpdatable.UpdateEpoch(currentEpoch)
	eu.epochUpdatables = append(eu.epochUpdatables, &epochUpdatable)
}

func (eu *EpochUpdater) UpdaterKey() string {
	return CallbackKeyForEpochUpdate
}

func (eu *EpochUpdater) Update(latestBlock int64) {
	eu.lock.RLock()
	defer eu.lock.RUnlock()
	ctx := context.Background()

	currentEpoch, err := eu.stateQuery.CurrentEpochStart(ctx)
	if err != nil {
		return // failed to get the current epoch
	}

	if currentEpoch <= eu.currentEpoch {
		return // still the same epoch
	}
	eu.currentEpoch = currentEpoch
	for _, epochUpdatable := range eu.epochUpdatables {
		if epochUpdatable == nil {
			continue
		}
		(*epochUpdatable).UpdateEpoch(currentEpoch)
	}
}

func (eu *EpochUpdater) EmergencyModeUpdate(virtualEpoch uint64) {
	eu.lock.RLock()
	defer eu.lock.RUnlock()

	if virtualEpoch <= eu.currentVirtualEpoch {
		return
	}

	utils.LavaFormatDebug("Emergency mode is turn on", utils.Attribute{Key: "virtual_epoch", Value: virtualEpoch})
	for _, epochUpdatable := range eu.epochUpdatables {
		if epochUpdatable == nil {
			continue
		}
		(*epochUpdatable).UpdateVirtualEpoch(eu.currentEpoch, virtualEpoch)
	}

	atomic.StoreUint64(&eu.currentVirtualEpoch, virtualEpoch)
}
