package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecuritySamlInvalidateFunc(t Transport) SecuritySamlInvalidate {
	_ = "STUB: not implemented"
	return *new(SecuritySamlInvalidate)
}

type SecuritySamlInvalidate func(body io.Reader, o ...func(*SecuritySamlInvalidateRequest)) (*Response, error)

type SecuritySamlInvalidateRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecuritySamlInvalidateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecuritySamlInvalidate) WithContext(v context.Context) func(*SecuritySamlInvalidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlInvalidate) WithPretty() func(*SecuritySamlInvalidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlInvalidate) WithHuman() func(*SecuritySamlInvalidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlInvalidate) WithErrorTrace() func(*SecuritySamlInvalidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlInvalidate) WithFilterPath(v ...string) func(*SecuritySamlInvalidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlInvalidate) WithHeader(h map[string]string) func(*SecuritySamlInvalidateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlInvalidate) WithOpaqueID(s string) func(*SecuritySamlInvalidateRequest) {
	_ = "STUB: not implemented"
	return nil
}
