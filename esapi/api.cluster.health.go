package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterHealthFunc(t Transport) ClusterHealth {
	_ = "STUB: not implemented"
	return *new(ClusterHealth)
}

type ClusterHealth func(o ...func(*ClusterHealthRequest)) (*Response, error)

type ClusterHealthRequest struct {
	Index []string

	ExpandWildcards             []string
	Level                       string
	Local                       *bool
	MasterTimeout               time.Duration
	Timeout                     time.Duration
	WaitForActiveShards         string
	WaitForEvents               string
	WaitForNoInitializingShards *bool
	WaitForNoRelocatingShards   *bool
	WaitForNodes                string
	WaitForStatus               string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterHealthRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterHealth) WithContext(v context.Context) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithIndex(v ...string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithExpandWildcards(v ...string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithLevel(v string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithLocal(v bool) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithMasterTimeout(v time.Duration) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithTimeout(v time.Duration) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithWaitForActiveShards(v string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithWaitForEvents(v string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithWaitForNoInitializingShards(v bool) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithWaitForNoRelocatingShards(v bool) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithWaitForNodes(v string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithWaitForStatus(v string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithPretty() func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithHuman() func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithErrorTrace() func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithFilterPath(v ...string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithHeader(h map[string]string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterHealth) WithOpaqueID(s string) func(*ClusterHealthRequest) {
	_ = "STUB: not implemented"
	return nil
}
