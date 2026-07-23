package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLValidateFunc(t Transport) MLValidate { _ = "STUB: not implemented"; return *new(MLValidate) }

type MLValidate func(body io.Reader, o ...func(*MLValidateRequest)) (*Response, error)

type MLValidateRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLValidateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLValidate) WithContext(v context.Context) func(*MLValidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidate) WithPretty() func(*MLValidateRequest) { _ = "STUB: not implemented"; return nil }

func (f MLValidate) WithHuman() func(*MLValidateRequest) { _ = "STUB: not implemented"; return nil }

func (f MLValidate) WithErrorTrace() func(*MLValidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidate) WithFilterPath(v ...string) func(*MLValidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidate) WithHeader(h map[string]string) func(*MLValidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidate) WithOpaqueID(s string) func(*MLValidateRequest) {
	_ = "STUB: not implemented"
	return nil
}
