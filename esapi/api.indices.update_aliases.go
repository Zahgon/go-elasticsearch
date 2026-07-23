package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesUpdateAliasesFunc(t Transport) IndicesUpdateAliases {
	_ = "STUB: not implemented"
	return *new(IndicesUpdateAliases)
}

type IndicesUpdateAliases func(body io.Reader, o ...func(*IndicesUpdateAliasesRequest)) (*Response, error)

type IndicesUpdateAliasesRequest struct {
	Body io.Reader

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesUpdateAliasesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesUpdateAliases) WithContext(v context.Context) func(*IndicesUpdateAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesUpdateAliases) WithMasterTimeout(v time.Duration) func(*IndicesUpdateAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesUpdateAliases) WithTimeout(v time.Duration) func(*IndicesUpdateAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesUpdateAliases) WithPretty() func(*IndicesUpdateAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesUpdateAliases) WithHuman() func(*IndicesUpdateAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesUpdateAliases) WithErrorTrace() func(*IndicesUpdateAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesUpdateAliases) WithFilterPath(v ...string) func(*IndicesUpdateAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesUpdateAliases) WithHeader(h map[string]string) func(*IndicesUpdateAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesUpdateAliases) WithOpaqueID(s string) func(*IndicesUpdateAliasesRequest) {
	_ = "STUB: not implemented"
	return nil
}
