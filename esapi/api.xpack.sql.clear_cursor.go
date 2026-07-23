package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSQLClearCursorFunc(t Transport) SQLClearCursor {
	_ = "STUB: not implemented"
	return *new(SQLClearCursor)
}

type SQLClearCursor func(body io.Reader, o ...func(*SQLClearCursorRequest)) (*Response, error)

type SQLClearCursorRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SQLClearCursorRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SQLClearCursor) WithContext(v context.Context) func(*SQLClearCursorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLClearCursor) WithPretty() func(*SQLClearCursorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLClearCursor) WithHuman() func(*SQLClearCursorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLClearCursor) WithErrorTrace() func(*SQLClearCursorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLClearCursor) WithFilterPath(v ...string) func(*SQLClearCursorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLClearCursor) WithHeader(h map[string]string) func(*SQLClearCursorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLClearCursor) WithOpaqueID(s string) func(*SQLClearCursorRequest) {
	_ = "STUB: not implemented"
	return nil
}
