package esapi

import (
	"context"
	"net/http"
)

func newClusterRemoteInfoFunc(t Transport) ClusterRemoteInfo {
	_ = "STUB: not implemented"
	return *new(ClusterRemoteInfo)
}

type ClusterRemoteInfo func(o ...func(*ClusterRemoteInfoRequest)) (*Response, error)

type ClusterRemoteInfoRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterRemoteInfoRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterRemoteInfo) WithContext(v context.Context) func(*ClusterRemoteInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterRemoteInfo) WithPretty() func(*ClusterRemoteInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterRemoteInfo) WithHuman() func(*ClusterRemoteInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterRemoteInfo) WithErrorTrace() func(*ClusterRemoteInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterRemoteInfo) WithFilterPath(v ...string) func(*ClusterRemoteInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterRemoteInfo) WithHeader(h map[string]string) func(*ClusterRemoteInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterRemoteInfo) WithOpaqueID(s string) func(*ClusterRemoteInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}
