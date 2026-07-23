package esapi

import (
	"context"
	"net/http"
)

func newMLGetFiltersFunc(t Transport) MLGetFilters {
	_ = "STUB: not implemented"
	return *new(MLGetFilters)
}

type MLGetFilters func(o ...func(*MLGetFiltersRequest)) (*Response, error)

type MLGetFiltersRequest struct {
	FilterID []string

	From *int
	Size *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetFiltersRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetFilters) WithContext(v context.Context) func(*MLGetFiltersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetFilters) WithFilterID(v ...string) func(*MLGetFiltersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetFilters) WithFrom(v int) func(*MLGetFiltersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetFilters) WithSize(v int) func(*MLGetFiltersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetFilters) WithPretty() func(*MLGetFiltersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetFilters) WithHuman() func(*MLGetFiltersRequest) { _ = "STUB: not implemented"; return nil }

func (f MLGetFilters) WithErrorTrace() func(*MLGetFiltersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetFilters) WithFilterPath(v ...string) func(*MLGetFiltersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetFilters) WithHeader(h map[string]string) func(*MLGetFiltersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetFilters) WithOpaqueID(s string) func(*MLGetFiltersRequest) {
	_ = "STUB: not implemented"
	return nil
}
