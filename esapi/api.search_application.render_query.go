package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSearchApplicationRenderQueryFunc(t Transport) SearchApplicationRenderQuery {
	_ = "STUB: not implemented"
	return *new(SearchApplicationRenderQuery)
}

type SearchApplicationRenderQuery func(name string, o ...func(*SearchApplicationRenderQueryRequest)) (*Response, error)

type SearchApplicationRenderQueryRequest struct {
	Body io.Reader

	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationRenderQueryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationRenderQuery) WithContext(v context.Context) func(*SearchApplicationRenderQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationRenderQuery) WithBody(v io.Reader) func(*SearchApplicationRenderQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationRenderQuery) WithPretty() func(*SearchApplicationRenderQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationRenderQuery) WithHuman() func(*SearchApplicationRenderQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationRenderQuery) WithErrorTrace() func(*SearchApplicationRenderQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationRenderQuery) WithFilterPath(v ...string) func(*SearchApplicationRenderQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationRenderQuery) WithHeader(h map[string]string) func(*SearchApplicationRenderQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationRenderQuery) WithOpaqueID(s string) func(*SearchApplicationRenderQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}
