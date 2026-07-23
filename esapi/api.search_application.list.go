package esapi

import (
	"context"
	"net/http"
)

func newSearchApplicationListFunc(t Transport) SearchApplicationList {
	_ = "STUB: not implemented"
	return *new(SearchApplicationList)
}

type SearchApplicationList func(o ...func(*SearchApplicationListRequest)) (*Response, error)

type SearchApplicationListRequest struct {
	From  *int
	Query string
	Size  *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationListRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationList) WithContext(v context.Context) func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationList) WithFrom(v int) func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationList) WithQuery(v string) func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationList) WithSize(v int) func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationList) WithPretty() func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationList) WithHuman() func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationList) WithErrorTrace() func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationList) WithFilterPath(v ...string) func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationList) WithHeader(h map[string]string) func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationList) WithOpaqueID(s string) func(*SearchApplicationListRequest) {
	_ = "STUB: not implemented"
	return nil
}
