package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesExistsAliasFunc(t Transport) IndicesExistsAlias {
	_ = "STUB: not implemented"
	return *new(IndicesExistsAlias)
}

type IndicesExistsAlias func(name []string, o ...func(*IndicesExistsAliasRequest)) (*Response, error)

type IndicesExistsAliasRequest struct {
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

func (r IndicesExistsAliasRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesExistsAlias) WithContext(v context.Context) func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithIndex(v ...string) func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithAllowNoIndices(v bool) func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithExpandWildcards(v ...string) func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithIgnoreUnavailable(v bool) func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithMasterTimeout(v time.Duration) func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithPretty() func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithHuman() func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithErrorTrace() func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithFilterPath(v ...string) func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithHeader(h map[string]string) func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExistsAlias) WithOpaqueID(s string) func(*IndicesExistsAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}
