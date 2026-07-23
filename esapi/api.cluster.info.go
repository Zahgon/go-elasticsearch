package esapi

import (
	"context"
	"net/http"
)

func newClusterInfoFunc(t Transport) ClusterInfo {
	_ = "STUB: not implemented"
	return *new(ClusterInfo)
}

type ClusterInfo func(target []string, o ...func(*ClusterInfoRequest)) (*Response, error)

type ClusterInfoRequest struct {
	Target []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterInfoRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterInfo) WithContext(v context.Context) func(*ClusterInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterInfo) WithPretty() func(*ClusterInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f ClusterInfo) WithHuman() func(*ClusterInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f ClusterInfo) WithErrorTrace() func(*ClusterInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterInfo) WithFilterPath(v ...string) func(*ClusterInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterInfo) WithHeader(h map[string]string) func(*ClusterInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterInfo) WithOpaqueID(s string) func(*ClusterInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}
