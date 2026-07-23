package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesPutAliasFunc(t Transport) IndicesPutAlias {
	_ = "STUB: not implemented"
	return *new(IndicesPutAlias)
}

type IndicesPutAlias func(index []string, name string, o ...func(*IndicesPutAliasRequest)) (*Response, error)

type IndicesPutAliasRequest struct {
	Index []string

	Body io.Reader

	Name string

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

func (r IndicesPutAliasRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPutAlias) WithContext(v context.Context) func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutAlias) WithBody(v io.Reader) func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutAlias) WithMasterTimeout(v time.Duration) func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutAlias) WithTimeout(v time.Duration) func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutAlias) WithPretty() func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutAlias) WithHuman() func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutAlias) WithErrorTrace() func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutAlias) WithFilterPath(v ...string) func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutAlias) WithHeader(h map[string]string) func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutAlias) WithOpaqueID(s string) func(*IndicesPutAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}
