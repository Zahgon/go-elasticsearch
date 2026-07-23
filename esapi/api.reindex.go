package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newReindexFunc(t Transport) Reindex { _ = "STUB: not implemented"; return *new(Reindex) }

type Reindex func(body io.Reader, o ...func(*ReindexRequest)) (*Response, error)

type ReindexRequest struct {
	Body io.Reader

	MaxDocs             *int
	Refresh             *bool
	RequestsPerSecond   *int
	RequireAlias        *bool
	Scroll              time.Duration
	Slices              interface{}
	Timeout             time.Duration
	WaitForActiveShards string
	WaitForCompletion   *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ReindexRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Reindex) WithContext(v context.Context) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithMaxDocs(v int) func(*ReindexRequest) { _ = "STUB: not implemented"; return nil }

func (f Reindex) WithRefresh(v bool) func(*ReindexRequest) { _ = "STUB: not implemented"; return nil }

func (f Reindex) WithRequestsPerSecond(v int) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithRequireAlias(v bool) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithScroll(v time.Duration) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithSlices(v interface{}) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithTimeout(v time.Duration) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithWaitForActiveShards(v string) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithWaitForCompletion(v bool) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithPretty() func(*ReindexRequest) { _ = "STUB: not implemented"; return nil }

func (f Reindex) WithHuman() func(*ReindexRequest) { _ = "STUB: not implemented"; return nil }

func (f Reindex) WithErrorTrace() func(*ReindexRequest) { _ = "STUB: not implemented"; return nil }

func (f Reindex) WithFilterPath(v ...string) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithHeader(h map[string]string) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Reindex) WithOpaqueID(s string) func(*ReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}
