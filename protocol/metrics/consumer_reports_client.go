package metrics

import (
	"encoding/json"
	"strings"

	"github.com/lavanet/lava/utils"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
)

const (
	reportName = "report"
)

type ConsumerReportsClient struct {
	*QueueSender
}

func NewReportsRequest(provider string, errors []error, specId string) ReportsRequest {
	errorsStrings := []string{}
	for _, err := range errors {
		errorsStrings = append(errorsStrings, err.Error())
	}
	return ReportsRequest{
		Name:     reportName,
		Errors:   strings.Join(errorsStrings, ","),
		Provider: provider,
		SpecId:   specId,
	}
}

type ReportsRequest struct {
	Name     string `json:"name"`
	Errors   string `json:"errors"`
	Provider string `json:"provider"`
	SpecId   string `json:"spec_id"`
}

func (rr ReportsRequest) String() string {
	rr.Name = reportName
	bytes, err := json.Marshal(rr)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type Reporter interface {
	AppendReport(report ReportsRequest)
}

type ConflictContainer struct {
	Request pairingtypes.RelayRequest `json:"request"`
	Reply   pairingtypes.RelayReply   `json:"reply"`
}

type ConflictRequest struct {
	Name      string              `json:"name"`
	Conflicts []ConflictContainer `json:"conflicts"`
}

func (rr ConflictRequest) String() string {
	rr.Name = "conflict"
	bytes, err := json.Marshal(rr)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func NewConsumerReportsClient(endpointAddress string) *ConsumerReportsClient {
	if endpointAddress == "" {
		utils.LavaFormatInfo("Running with Consumer Relay Server Disabled")
		return nil
	}
	cuc := &ConsumerReportsClient{
		QueueSender: NewQueueSender(endpointAddress, "ConsumerReports"),
	}
	return cuc
}

func (cuc *ConsumerReportsClient) AppendReport(report ReportsRequest) {
	cuc.appendQueue(report)
}

func (cuc *ConsumerReportsClient) AppendConflict(report ReportsRequest) {
	cuc.appendQueue(report)
}
