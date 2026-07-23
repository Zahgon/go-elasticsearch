package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSQLTranslateFunc(t Transport) SQLTranslate {
	_ = "STUB: not implemented"
	return *new(SQLTranslate)
}

type SQLTranslate func(body io.Reader, o ...func(*SQLTranslateRequest)) (*Response, error)

type SQLTranslateRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SQLTranslateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SQLTranslate) WithContext(v context.Context) func(*SQLTranslateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLTranslate) WithPretty() func(*SQLTranslateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLTranslate) WithHuman() func(*SQLTranslateRequest) { _ = "STUB: not implemented"; return nil }

func (f SQLTranslate) WithErrorTrace() func(*SQLTranslateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLTranslate) WithFilterPath(v ...string) func(*SQLTranslateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLTranslate) WithHeader(h map[string]string) func(*SQLTranslateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLTranslate) WithOpaqueID(s string) func(*SQLTranslateRequest) {
	_ = "STUB: not implemented"
	return nil
}
