package esapi

import (
	"context"
	"net/http"
	"time"
)

func newNodesHotThreadsFunc(t Transport) NodesHotThreads {
	_ = "STUB: not implemented"
	return *new(NodesHotThreads)
}

type NodesHotThreads func(o ...func(*NodesHotThreadsRequest)) (*Response, error)

type NodesHotThreadsRequest struct {
	NodeID []string

	IgnoreIdleThreads *bool
	Interval          time.Duration
	Snapshots         *int64
	Sort              string
	Threads           *int64
	Timeout           time.Duration
	DocumentType      string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r NodesHotThreadsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f NodesHotThreads) WithContext(v context.Context) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithNodeID(v ...string) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithIgnoreIdleThreads(v bool) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithInterval(v time.Duration) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithSnapshots(v int64) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithSort(v string) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithThreads(v int64) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithTimeout(v time.Duration) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithDocumentType(v string) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithPretty() func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithHuman() func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithErrorTrace() func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithFilterPath(v ...string) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithHeader(h map[string]string) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesHotThreads) WithOpaqueID(s string) func(*NodesHotThreadsRequest) {
	_ = "STUB: not implemented"
	return nil
}
