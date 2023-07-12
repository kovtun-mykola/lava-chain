package chainlib

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/lavanet/lava/protocol/chainlib/chainproxy"
	"github.com/lavanet/lava/protocol/chainlib/chainproxy/rpcInterfaceMessages"
	"github.com/lavanet/lava/protocol/parser"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	connectionType_test = "test"
)

func TestGRPCChainParser_Spec(t *testing.T) {
	// create a new instance of RestChainParser
	apip, err := NewGrpcChainParser()
	if err != nil {
		t.Errorf("Error creating RestChainParser: %v", err)
	}

	// set the spec
	spec := spectypes.Spec{
		Enabled:                       true,
		ReliabilityThreshold:          10,
		AllowedBlockLagForQosSync:     11,
		AverageBlockTime:              12000,
		BlockDistanceForFinalizedData: 13,
		BlocksInFinalizationProof:     14,
	}
	apip.SetSpec(spec)

	// fetch data reliability params
	enabled, dataReliabilityThreshold := apip.DataReliabilityParams()

	// fetch chain block stats
	allowedBlockLagForQosSync, averageBlockTime, blockDistanceForFinalizedData, blocksInFinalizationProof := apip.ChainBlockStats()

	// convert block time
	AverageBlockTime := time.Duration(apip.spec.AverageBlockTime) * time.Millisecond

	// check that the spec was set correctly
	assert.Equal(t, apip.spec.DataReliabilityEnabled, enabled)
	assert.Equal(t, apip.spec.GetReliabilityThreshold(), dataReliabilityThreshold)
	assert.Equal(t, apip.spec.AllowedBlockLagForQosSync, allowedBlockLagForQosSync)
	assert.Equal(t, apip.spec.BlockDistanceForFinalizedData, blockDistanceForFinalizedData)
	assert.Equal(t, apip.spec.BlocksInFinalizationProof, blocksInFinalizationProof)
	assert.Equal(t, AverageBlockTime, averageBlockTime)
}

func TestGRPChainParser_NilGuard(t *testing.T) {
	var apip *GrpcChainParser

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("apip methods missing nill guard, panicked with: %v", r)
		}
	}()

	apip.SetSpec(spectypes.Spec{})
	apip.DataReliabilityParams()
	apip.ChainBlockStats()
	apip.getSupportedApi("", "")
	apip.ParseMsg("", []byte{}, "", nil)
}

func TestGRPCGetSupportedApi(t *testing.T) {
	// Test case 1: Successful scenario, returns a supported API
	apip := &GrpcChainParser{
		BaseChainParser: BaseChainParser{
			serverApis: map[ApiKey]ApiContainer{{Name: "API1", ConnectionType: connectionType_test}: {api: &spectypes.Api{Name: "API1", Enabled: true}, collectionKey: CollectionKey{ConnectionType: connectionType_test}}},
		},
	}
	apiCont, err := apip.getSupportedApi("API1", connectionType_test)
	assert.NoError(t, err)
	assert.Equal(t, "API1", apiCont.api.Name)

	// Test case 2: Returns error if the API does not exist
	apip = &GrpcChainParser{
		BaseChainParser: BaseChainParser{
			serverApis: map[ApiKey]ApiContainer{{Name: "API1", ConnectionType: connectionType_test}: {api: &spectypes.Api{Name: "API1", Enabled: true}, collectionKey: CollectionKey{ConnectionType: connectionType_test}}},
		},
	}
	_, err = apip.getSupportedApi("API2", connectionType_test)
	assert.Error(t, err)
	errorData, _, found := strings.Cut(err.Error(), " --")
	require.True(t, found)
	assert.Equal(t, "api not supported", errorData)

	// Test case 3: Returns error if the API is disabled
	apip = &GrpcChainParser{
		BaseChainParser: BaseChainParser{
			serverApis: map[ApiKey]ApiContainer{{Name: "API1", ConnectionType: connectionType_test}: {api: &spectypes.Api{Name: "API1", Enabled: false}, collectionKey: CollectionKey{ConnectionType: connectionType_test}}},
		},
	}
	_, err = apip.getSupportedApi("API1", connectionType_test)
	assert.Error(t, err)
	errorData, _, found = strings.Cut(err.Error(), " --")
	require.True(t, found)
	assert.Equal(t, "api is disabled", errorData)
}

func TestGRPCParseMessage(t *testing.T) {
	apip := &GrpcChainParser{
		BaseChainParser: BaseChainParser{
			serverApis: map[ApiKey]ApiContainer{
				{Name: "API1", ConnectionType: connectionType_test}: {api: &spectypes.Api{Name: "API1", Enabled: true}, collectionKey: CollectionKey{ConnectionType: connectionType_test}},
			},
			apiCollections: map[CollectionKey]*spectypes.ApiCollection{{ConnectionType: connectionType_test}: {Enabled: true, CollectionData: spectypes.CollectionData{ApiInterface: spectypes.APIInterfaceGrpc}}},
		},
	}

	msg, err := apip.ParseMsg("API1", []byte("test message"), connectionType_test, nil)

	assert.Nil(t, err)
	assert.Equal(t, msg.GetApi().Name, apip.serverApis[ApiKey{Name: "API1", ConnectionType: connectionType_test}].api.Name)
	assert.Equal(t, msg.GetApiCollection().CollectionData.ApiInterface, spectypes.APIInterfaceGrpc)

	grpcMessage := rpcInterfaceMessages.GrpcMessage{
		Msg:         []byte("test message"),
		Path:        "API1",
		BaseMessage: chainproxy.BaseMessage{Headers: []pairingtypes.Metadata{}},
	}
	grpcMsg, ok := msg.GetRPCMessage().(*rpcInterfaceMessages.GrpcMessage)
	require.True(t, ok)
	assert.Equal(t, grpcMessage, *grpcMsg)
}

func TestGrpcChainProxy(t *testing.T) {
	ctx := context.Background()
	wasCalled := false
	serverHandle := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle the incoming request and provide the desired response
		wasCalled = true
	})
	chainParser, chainProxy, chainFetcher, closeServer, err := CreateChainLibMocks(ctx, "LAV1", spectypes.APIInterfaceGrpc, serverHandle, "../../")
	require.NoError(t, err)
	require.NotNil(t, chainParser)
	require.NotNil(t, chainProxy)
	require.NotNil(t, chainFetcher)
	_, err = chainFetcher.FetchLatestBlockNum(ctx)
	require.True(t, wasCalled)
	require.NoError(t, err)
	if closeServer != nil {
		closeServer()
	}
}

func TestParsingRequestedBlocksHeadersGrpc(t *testing.T) {
	ctx := context.Background()
	callbackHeaderNameToCheck := ""
	serverHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle the incoming request and provide the desired response
		headers := r.Header[callbackHeaderNameToCheck]
		if len(headers) > 0 && headers[0] != "" {
			block, err := strconv.Atoi(headers[0])
			require.NoError(t, err)
			w.WriteHeader(block)
		} else {
			w.WriteHeader(244591)
		}
	})
	chainParser, chainProxy, _, closeServer, err := CreateChainLibMocks(ctx, "LAV1", spectypes.APIInterfaceGrpc, serverHandler, "../../")
	require.NoError(t, err)
	defer func() {
		if closeServer != nil {
			closeServer()
		}
	}()
	parsingForCrafting, collectionData, ok := chainParser.GetParsingByTag(spectypes.FUNCTION_TAG_GET_BLOCKNUM)
	require.True(t, ok)
	headerParsingDirective, _, ok := chainParser.GetParsingByTag(spectypes.FUNCTION_TAG_SET_LATEST_IN_METADATA)
	callbackHeaderNameToCheck = headerParsingDirective.GetApiName() // this causes the callback to modify the response to simulate a real behavior
	require.True(t, ok)
	block := 244590
	metadata := []pairingtypes.Metadata{{Name: headerParsingDirective.GetApiName(), Value: fmt.Sprintf(headerParsingDirective.FunctionTemplate, block)}}

	tests := []struct {
		desc           string
		metadata       []pairingtypes.Metadata
		block          int64
		requestedBlock int64
	}{
		{
			desc:           "no metadata",
			metadata:       []pairingtypes.Metadata{},
			block:          244591,
			requestedBlock: spectypes.LATEST_BLOCK,
		},
		{
			desc:           "with-metadata",
			metadata:       metadata,
			block:          244590,
			requestedBlock: 244590,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainMessage, err := chainParser.ParseMsg(parsingForCrafting.ApiName, []byte{}, collectionData.Type, test.metadata)
			require.NoError(t, err)
			require.NoError(t, err)

			require.Equal(t, test.requestedBlock, chainMessage.RequestedBlock())
			reply, _, _, err := chainProxy.SendNodeMsg(ctx, nil, chainMessage)
			require.NoError(t, err)
			parserInput, err := FormatResponseForParsing(reply, chainMessage)
			require.NoError(t, err)
			blockNum, err := parser.ParseBlockFromReply(parserInput, parsingForCrafting.ResultParsing)
			require.NoError(t, err)
			require.Equal(t, test.block, blockNum)
		})
	}
}

func TestSettingBlocksHeadersGrpc(t *testing.T) {
	ctx := context.Background()
	callbackHeaderNameToCheck := ""
	serverHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle the incoming request and provide the desired response
		headers := r.Header[callbackHeaderNameToCheck]
		if len(headers) > 0 && headers[0] != "" {
			block, err := strconv.Atoi(headers[0])
			require.NoError(t, err)
			w.WriteHeader(block)
		} else {
			w.WriteHeader(244591)
		}
	})
	chainParser, chainProxy, _, closeServer, err := CreateChainLibMocks(ctx, "LAV1", spectypes.APIInterfaceGrpc, serverHandler, "../../")
	require.NoError(t, err)
	defer func() {
		if closeServer != nil {
			closeServer()
		}
	}()
	parsingForCrafting, collectionData, ok := chainParser.GetParsingByTag(spectypes.FUNCTION_TAG_GET_BLOCKNUM)
	require.True(t, ok)
	headerParsingDirective, _, ok := chainParser.GetParsingByTag(spectypes.FUNCTION_TAG_SET_LATEST_IN_METADATA)
	callbackHeaderNameToCheck = headerParsingDirective.GetApiName() // this causes the callback to modify the response to simulate a real behavior
	require.True(t, ok)
	block := 244590
	metadata := []pairingtypes.Metadata{{Name: headerParsingDirective.GetApiName(), Value: fmt.Sprintf(headerParsingDirective.FunctionTemplate, block)}}

	tests := []struct {
		desc           string
		metadata       []pairingtypes.Metadata
		block          int64
		requestedBlock int64
	}{
		{
			desc:           "no metadata",
			metadata:       []pairingtypes.Metadata{},
			block:          244589,
			requestedBlock: spectypes.LATEST_BLOCK,
		},
		{
			desc:           "with-metadata",
			metadata:       metadata,
			block:          244590,
			requestedBlock: 244590,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainMessage, err := chainParser.ParseMsg(parsingForCrafting.ApiName, []byte{}, collectionData.Type, test.metadata)
			require.NoError(t, err)
			require.Equal(t, test.requestedBlock, chainMessage.RequestedBlock())
			chainMessage.UpdateLatestBlockInMessage(test.block, true) // will update the request only if it's latest
			require.Equal(t, test.block, chainMessage.RequestedBlock())
			reply, _, _, err := chainProxy.SendNodeMsg(ctx, nil, chainMessage)
			require.NoError(t, err)
			parserInput, err := FormatResponseForParsing(reply, chainMessage)
			require.NoError(t, err)
			blockNum, err := parser.ParseBlockFromReply(parserInput, parsingForCrafting.ResultParsing)
			require.NoError(t, err)
			require.Equal(t, test.block, blockNum)
		})
	}
}
