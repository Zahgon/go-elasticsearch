package esapi

import (
	"context"
	"net/http"
)

func newIndicesShardStoresFunc(t Transport) IndicesShardStores {
	_ = "STUB: not implemented"
	return *new(IndicesShardStores)
}

type IndicesShardStores func(o ...func(*IndicesShardStoresRequest)) (*Response, error)

type IndicesShardStoresRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	Status            []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesShardStoresRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesShardStores) WithContext(v context.Context) func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithIndex(v ...string) func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithAllowNoIndices(v bool) func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithExpandWildcards(v ...string) func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithIgnoreUnavailable(v bool) func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithStatus(v ...string) func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithPretty() func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithHuman() func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithErrorTrace() func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithFilterPath(v ...string) func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithHeader(h map[string]string) func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShardStores) WithOpaqueID(s string) func(*IndicesShardStoresRequest) {
	_ = "STUB: not implemented"
	return nil
}
