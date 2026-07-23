package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetAliasFunc(t Transport) IndicesGetAlias {
	_ = "STUB: not implemented"
	return *new(IndicesGetAlias)
}

type IndicesGetAlias func(o ...func(*IndicesGetAliasRequest)) (*Response, error)

type IndicesGetAliasRequest struct {
	Index []string

	Name []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	MasterTimeout     time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetAliasRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetAlias) WithContext(v context.Context) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithIndex(v ...string) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithName(v ...string) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithAllowNoIndices(v bool) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithExpandWildcards(v ...string) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithIgnoreUnavailable(v bool) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithMasterTimeout(v time.Duration) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithPretty() func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithHuman() func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithErrorTrace() func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithFilterPath(v ...string) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithHeader(h map[string]string) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetAlias) WithOpaqueID(s string) func(*IndicesGetAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}
