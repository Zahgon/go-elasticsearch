package esapi

import (
	"context"
	"net/http"
)

func newIndicesForcemergeFunc(t Transport) IndicesForcemerge {
	_ = "STUB: not implemented"
	return *new(IndicesForcemerge)
}

type IndicesForcemerge func(o ...func(*IndicesForcemergeRequest)) (*Response, error)

type IndicesForcemergeRequest struct {
	Index []string

	AllowNoIndices     *bool
	ExpandWildcards    []string
	Flush              *bool
	IgnoreUnavailable  *bool
	MaxNumSegments     *int64
	OnlyExpungeDeletes *bool
	WaitForCompletion  *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesForcemergeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesForcemerge) WithContext(v context.Context) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithIndex(v ...string) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithAllowNoIndices(v bool) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithExpandWildcards(v ...string) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithFlush(v bool) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithIgnoreUnavailable(v bool) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithMaxNumSegments(v int64) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithOnlyExpungeDeletes(v bool) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithWaitForCompletion(v bool) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithPretty() func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithHuman() func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithErrorTrace() func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithFilterPath(v ...string) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithHeader(h map[string]string) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesForcemerge) WithOpaqueID(s string) func(*IndicesForcemergeRequest) {
	_ = "STUB: not implemented"
	return nil
}
