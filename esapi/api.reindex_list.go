package esapi

import (
	"context"
	"net/http"
)

func newReindexListFunc(t Transport) ReindexList {
	_ = "STUB: not implemented"
	return *new(ReindexList)
}

type ReindexList func(o ...func(*ReindexListRequest)) (*Response, error)

type ReindexListRequest struct {
	Detailed *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ReindexListRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ReindexList) WithContext(v context.Context) func(*ReindexListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexList) WithDetailed(v bool) func(*ReindexListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexList) WithPretty() func(*ReindexListRequest) { _ = "STUB: not implemented"; return nil }

func (f ReindexList) WithHuman() func(*ReindexListRequest) { _ = "STUB: not implemented"; return nil }

func (f ReindexList) WithErrorTrace() func(*ReindexListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexList) WithFilterPath(v ...string) func(*ReindexListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexList) WithHeader(h map[string]string) func(*ReindexListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexList) WithOpaqueID(s string) func(*ReindexListRequest) {
	_ = "STUB: not implemented"
	return nil
}
