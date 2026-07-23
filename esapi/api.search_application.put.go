package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSearchApplicationPutFunc(t Transport) SearchApplicationPut {
	_ = "STUB: not implemented"
	return *new(SearchApplicationPut)
}

type SearchApplicationPut func(name string, body io.Reader, o ...func(*SearchApplicationPutRequest)) (*Response, error)

type SearchApplicationPutRequest struct {
	Body io.Reader

	Name string

	Create *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationPutRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationPut) WithContext(v context.Context) func(*SearchApplicationPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPut) WithCreate(v bool) func(*SearchApplicationPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPut) WithPretty() func(*SearchApplicationPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPut) WithHuman() func(*SearchApplicationPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPut) WithErrorTrace() func(*SearchApplicationPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPut) WithFilterPath(v ...string) func(*SearchApplicationPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPut) WithHeader(h map[string]string) func(*SearchApplicationPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPut) WithOpaqueID(s string) func(*SearchApplicationPutRequest) {
	_ = "STUB: not implemented"
	return nil
}
