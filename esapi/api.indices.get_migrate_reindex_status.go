package esapi

import (
	"context"
	"net/http"
)

func newIndicesGetMigrateReindexStatusFunc(t Transport) IndicesGetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return *new(IndicesGetMigrateReindexStatus)
}

type IndicesGetMigrateReindexStatus func(index []string, o ...func(*IndicesGetMigrateReindexStatusRequest)) (*Response, error)

type IndicesGetMigrateReindexStatusRequest struct {
	Index []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetMigrateReindexStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetMigrateReindexStatus) WithContext(v context.Context) func(*IndicesGetMigrateReindexStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMigrateReindexStatus) WithPretty() func(*IndicesGetMigrateReindexStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMigrateReindexStatus) WithHuman() func(*IndicesGetMigrateReindexStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMigrateReindexStatus) WithErrorTrace() func(*IndicesGetMigrateReindexStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMigrateReindexStatus) WithFilterPath(v ...string) func(*IndicesGetMigrateReindexStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMigrateReindexStatus) WithHeader(h map[string]string) func(*IndicesGetMigrateReindexStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMigrateReindexStatus) WithOpaqueID(s string) func(*IndicesGetMigrateReindexStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
