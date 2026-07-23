package esapi

import (
	"context"
	"net/http"
	"time"
)

func newNodesUsageFunc(t Transport) NodesUsage { _ = "STUB: not implemented"; return *new(NodesUsage) }

type NodesUsage func(o ...func(*NodesUsageRequest)) (*Response, error)

type NodesUsageRequest struct {
	Metric []string
	NodeID []string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r NodesUsageRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f NodesUsage) WithContext(v context.Context) func(*NodesUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesUsage) WithMetric(v ...string) func(*NodesUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesUsage) WithNodeID(v ...string) func(*NodesUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesUsage) WithTimeout(v time.Duration) func(*NodesUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesUsage) WithPretty() func(*NodesUsageRequest) { _ = "STUB: not implemented"; return nil }

func (f NodesUsage) WithHuman() func(*NodesUsageRequest) { _ = "STUB: not implemented"; return nil }

func (f NodesUsage) WithErrorTrace() func(*NodesUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesUsage) WithFilterPath(v ...string) func(*NodesUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesUsage) WithHeader(h map[string]string) func(*NodesUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesUsage) WithOpaqueID(s string) func(*NodesUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}
