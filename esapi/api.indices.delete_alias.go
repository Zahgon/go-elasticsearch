package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesDeleteAliasFunc(t Transport) IndicesDeleteAlias {
	_ = "STUB: not implemented"
	return *new(IndicesDeleteAlias)
}

type IndicesDeleteAlias func(index []string, name []string, o ...func(*IndicesDeleteAliasRequest)) (*Response, error)

type IndicesDeleteAliasRequest struct {
	Index []string

	Name []string

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

func (r IndicesDeleteAliasRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDeleteAlias) WithContext(v context.Context) func(*IndicesDeleteAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteAlias) WithMasterTimeout(v time.Duration) func(*IndicesDeleteAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteAlias) WithTimeout(v time.Duration) func(*IndicesDeleteAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteAlias) WithPretty() func(*IndicesDeleteAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteAlias) WithHuman() func(*IndicesDeleteAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteAlias) WithErrorTrace() func(*IndicesDeleteAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteAlias) WithFilterPath(v ...string) func(*IndicesDeleteAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteAlias) WithHeader(h map[string]string) func(*IndicesDeleteAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteAlias) WithOpaqueID(s string) func(*IndicesDeleteAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}
