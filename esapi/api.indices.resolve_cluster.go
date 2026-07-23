package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesResolveClusterFunc(t Transport) IndicesResolveCluster {
	_ = "STUB: not implemented"
	return *new(IndicesResolveCluster)
}

type IndicesResolveCluster func(o ...func(*IndicesResolveClusterRequest)) (*Response, error)

type IndicesResolveClusterRequest struct {
	Name []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreThrottled   *bool
	IgnoreUnavailable *bool
	Timeout           time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesResolveClusterRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesResolveCluster) WithContext(v context.Context) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithName(v ...string) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithAllowNoIndices(v bool) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithExpandWildcards(v ...string) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithIgnoreThrottled(v bool) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithIgnoreUnavailable(v bool) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithTimeout(v time.Duration) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithPretty() func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithHuman() func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithErrorTrace() func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithFilterPath(v ...string) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithHeader(h map[string]string) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveCluster) WithOpaqueID(s string) func(*IndicesResolveClusterRequest) {
	_ = "STUB: not implemented"
	return nil
}
