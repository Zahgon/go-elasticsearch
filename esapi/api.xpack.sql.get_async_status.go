package esapi

import (
	"context"
	"net/http"
)

func newSQLGetAsyncStatusFunc(t Transport) SQLGetAsyncStatus {
	_ = "STUB: not implemented"
	return *new(SQLGetAsyncStatus)
}

type SQLGetAsyncStatus func(id string, o ...func(*SQLGetAsyncStatusRequest)) (*Response, error)

type SQLGetAsyncStatusRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SQLGetAsyncStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SQLGetAsyncStatus) WithContext(v context.Context) func(*SQLGetAsyncStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsyncStatus) WithPretty() func(*SQLGetAsyncStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsyncStatus) WithHuman() func(*SQLGetAsyncStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsyncStatus) WithErrorTrace() func(*SQLGetAsyncStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsyncStatus) WithFilterPath(v ...string) func(*SQLGetAsyncStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsyncStatus) WithHeader(h map[string]string) func(*SQLGetAsyncStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsyncStatus) WithOpaqueID(s string) func(*SQLGetAsyncStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
