package esapi

import (
	"context"
	"net/http"
	"time"
)

func newNodesInfoFunc(t Transport) NodesInfo { _ = "STUB: not implemented"; return *new(NodesInfo) }

type NodesInfo func(o ...func(*NodesInfoRequest)) (*Response, error)

type NodesInfoRequest struct {
	Metric []string
	NodeID []string

	FlatSettings *bool
	Timeout      time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r NodesInfoRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f NodesInfo) WithContext(v context.Context) func(*NodesInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesInfo) WithMetric(v ...string) func(*NodesInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesInfo) WithNodeID(v ...string) func(*NodesInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesInfo) WithFlatSettings(v bool) func(*NodesInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesInfo) WithTimeout(v time.Duration) func(*NodesInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesInfo) WithPretty() func(*NodesInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f NodesInfo) WithHuman() func(*NodesInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f NodesInfo) WithErrorTrace() func(*NodesInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f NodesInfo) WithFilterPath(v ...string) func(*NodesInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesInfo) WithHeader(h map[string]string) func(*NodesInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesInfo) WithOpaqueID(s string) func(*NodesInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}
