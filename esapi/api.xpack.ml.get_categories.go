package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLGetCategoriesFunc(t Transport) MLGetCategories {
	_ = "STUB: not implemented"
	return *new(MLGetCategories)
}

type MLGetCategories func(job_id string, o ...func(*MLGetCategoriesRequest)) (*Response, error)

type MLGetCategoriesRequest struct {
	Body io.Reader

	CategoryID *int64
	JobID      string

	From                *int
	PartitionFieldValue string
	Size                *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetCategoriesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetCategories) WithContext(v context.Context) func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithBody(v io.Reader) func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithCategoryID(v int64) func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithFrom(v int) func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithPartitionFieldValue(v string) func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithSize(v int) func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithPretty() func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithHuman() func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithErrorTrace() func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithFilterPath(v ...string) func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithHeader(h map[string]string) func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCategories) WithOpaqueID(s string) func(*MLGetCategoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}
