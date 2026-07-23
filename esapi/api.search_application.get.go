package esapi

import (
	"context"
	"net/http"
)

func newSearchApplicationGetFunc(t Transport) SearchApplicationGet {
	_ = "STUB: not implemented"
	return *new(SearchApplicationGet)
}

type SearchApplicationGet func(name string, o ...func(*SearchApplicationGetRequest)) (*Response, error)

type SearchApplicationGetRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationGet) WithContext(v context.Context) func(*SearchApplicationGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGet) WithPretty() func(*SearchApplicationGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGet) WithHuman() func(*SearchApplicationGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGet) WithErrorTrace() func(*SearchApplicationGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGet) WithFilterPath(v ...string) func(*SearchApplicationGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGet) WithHeader(h map[string]string) func(*SearchApplicationGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGet) WithOpaqueID(s string) func(*SearchApplicationGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
