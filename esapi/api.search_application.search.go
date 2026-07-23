package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSearchApplicationSearchFunc(t Transport) SearchApplicationSearch {
	_ = "STUB: not implemented"
	return *new(SearchApplicationSearch)
}

type SearchApplicationSearch func(name string, o ...func(*SearchApplicationSearchRequest)) (*Response, error)

type SearchApplicationSearchRequest struct {
	Body io.Reader

	Name string

	TypedKeys *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationSearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationSearch) WithContext(v context.Context) func(*SearchApplicationSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationSearch) WithBody(v io.Reader) func(*SearchApplicationSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationSearch) WithTypedKeys(v bool) func(*SearchApplicationSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationSearch) WithPretty() func(*SearchApplicationSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationSearch) WithHuman() func(*SearchApplicationSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationSearch) WithErrorTrace() func(*SearchApplicationSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationSearch) WithFilterPath(v ...string) func(*SearchApplicationSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationSearch) WithHeader(h map[string]string) func(*SearchApplicationSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationSearch) WithOpaqueID(s string) func(*SearchApplicationSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}
