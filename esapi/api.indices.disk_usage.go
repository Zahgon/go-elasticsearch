package esapi

import (
	"context"
	"net/http"
)

func newIndicesDiskUsageFunc(t Transport) IndicesDiskUsage {
	_ = "STUB: not implemented"
	return *new(IndicesDiskUsage)
}

type IndicesDiskUsage func(index []string, o ...func(*IndicesDiskUsageRequest)) (*Response, error)

type IndicesDiskUsageRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	Flush             *bool
	IgnoreUnavailable *bool
	RunExpensiveTasks *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesDiskUsageRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDiskUsage) WithContext(v context.Context) func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithAllowNoIndices(v bool) func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithExpandWildcards(v ...string) func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithFlush(v bool) func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithIgnoreUnavailable(v bool) func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithRunExpensiveTasks(v bool) func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithPretty() func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithHuman() func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithErrorTrace() func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithFilterPath(v ...string) func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithHeader(h map[string]string) func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDiskUsage) WithOpaqueID(s string) func(*IndicesDiskUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}
