package esapi

import (
	"context"
	"net/http"
)

func newIndicesClearCacheFunc(t Transport) IndicesClearCache {
	_ = "STUB: not implemented"
	return *new(IndicesClearCache)
}

type IndicesClearCache func(o ...func(*IndicesClearCacheRequest)) (*Response, error)

type IndicesClearCacheRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	Fielddata         *bool
	Fields            []string
	IgnoreUnavailable *bool
	Query             *bool
	Request           *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesClearCacheRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesClearCache) WithContext(v context.Context) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithIndex(v ...string) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithAllowNoIndices(v bool) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithExpandWildcards(v ...string) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithFielddata(v bool) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithFields(v ...string) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithIgnoreUnavailable(v bool) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithQuery(v bool) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithRequest(v bool) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithPretty() func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithHuman() func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithErrorTrace() func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithFilterPath(v ...string) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithHeader(h map[string]string) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClearCache) WithOpaqueID(s string) func(*IndicesClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}
