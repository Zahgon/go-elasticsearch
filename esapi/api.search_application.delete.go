package esapi

import (
	"context"
	"net/http"
)

func newSearchApplicationDeleteFunc(t Transport) SearchApplicationDelete {
	_ = "STUB: not implemented"
	return *new(SearchApplicationDelete)
}

type SearchApplicationDelete func(name string, o ...func(*SearchApplicationDeleteRequest)) (*Response, error)

type SearchApplicationDeleteRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationDelete) WithContext(v context.Context) func(*SearchApplicationDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDelete) WithPretty() func(*SearchApplicationDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDelete) WithHuman() func(*SearchApplicationDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDelete) WithErrorTrace() func(*SearchApplicationDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDelete) WithFilterPath(v ...string) func(*SearchApplicationDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDelete) WithHeader(h map[string]string) func(*SearchApplicationDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDelete) WithOpaqueID(s string) func(*SearchApplicationDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
