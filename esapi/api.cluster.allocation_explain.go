package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newClusterAllocationExplainFunc(t Transport) ClusterAllocationExplain {
	_ = "STUB: not implemented"
	return *new(ClusterAllocationExplain)
}

type ClusterAllocationExplain func(o ...func(*ClusterAllocationExplainRequest)) (*Response, error)

type ClusterAllocationExplainRequest struct {
	Body io.Reader

	CurrentNode         string
	IncludeDiskInfo     *bool
	IncludeYesDecisions *bool
	Index               string
	MasterTimeout       time.Duration
	Primary             *bool
	Shard               *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterAllocationExplainRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterAllocationExplain) WithContext(v context.Context) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithBody(v io.Reader) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithCurrentNode(v string) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithIncludeDiskInfo(v bool) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithIncludeYesDecisions(v bool) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithIndex(v string) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithMasterTimeout(v time.Duration) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithPrimary(v bool) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithShard(v int) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithPretty() func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithHuman() func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithErrorTrace() func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithFilterPath(v ...string) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithHeader(h map[string]string) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterAllocationExplain) WithOpaqueID(s string) func(*ClusterAllocationExplainRequest) {
	_ = "STUB: not implemented"
	return nil
}
