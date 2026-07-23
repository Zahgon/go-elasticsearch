package esapi

import (
	"context"
	"net/http"
)

func newIndicesFlushFunc(t Transport) IndicesFlush {
	_ = "STUB: not implemented"
	return *new(IndicesFlush)
}

type IndicesFlush func(o ...func(*IndicesFlushRequest)) (*Response, error)

type IndicesFlushRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	Force             *bool
	IgnoreUnavailable *bool
	WaitIfOngoing     *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesFlushRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesFlush) WithContext(v context.Context) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithIndex(v ...string) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithAllowNoIndices(v bool) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithExpandWildcards(v ...string) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithForce(v bool) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithIgnoreUnavailable(v bool) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithWaitIfOngoing(v bool) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithPretty() func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithHuman() func(*IndicesFlushRequest) { _ = "STUB: not implemented"; return nil }

func (f IndicesFlush) WithErrorTrace() func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithFilterPath(v ...string) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithHeader(h map[string]string) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFlush) WithOpaqueID(s string) func(*IndicesFlushRequest) {
	_ = "STUB: not implemented"
	return nil
}
