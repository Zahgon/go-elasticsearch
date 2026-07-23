package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterStateFunc(t Transport) ClusterState {
	_ = "STUB: not implemented"
	return *new(ClusterState)
}

type ClusterState func(o ...func(*ClusterStateRequest)) (*Response, error)

type ClusterStateRequest struct {
	Index []string

	Metric []string

	AllowNoIndices         *bool
	ExpandWildcards        []string
	FlatSettings           *bool
	IgnoreUnavailable      *bool
	Local                  *bool
	MasterTimeout          time.Duration
	WaitForMetadataVersion *int64
	WaitForTimeout         time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterStateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterState) WithContext(v context.Context) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithIndex(v ...string) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithMetric(v ...string) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithAllowNoIndices(v bool) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithExpandWildcards(v ...string) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithFlatSettings(v bool) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithIgnoreUnavailable(v bool) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithLocal(v bool) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithMasterTimeout(v time.Duration) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithWaitForMetadataVersion(v int64) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithWaitForTimeout(v time.Duration) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithPretty() func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithHuman() func(*ClusterStateRequest) { _ = "STUB: not implemented"; return nil }

func (f ClusterState) WithErrorTrace() func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithFilterPath(v ...string) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithHeader(h map[string]string) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterState) WithOpaqueID(s string) func(*ClusterStateRequest) {
	_ = "STUB: not implemented"
	return nil
}
