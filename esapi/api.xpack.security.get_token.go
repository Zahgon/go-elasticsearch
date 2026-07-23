package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityGetTokenFunc(t Transport) SecurityGetToken {
	_ = "STUB: not implemented"
	return *new(SecurityGetToken)
}

type SecurityGetToken func(body io.Reader, o ...func(*SecurityGetTokenRequest)) (*Response, error)

type SecurityGetTokenRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetTokenRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetToken) WithContext(v context.Context) func(*SecurityGetTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetToken) WithPretty() func(*SecurityGetTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetToken) WithHuman() func(*SecurityGetTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetToken) WithErrorTrace() func(*SecurityGetTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetToken) WithFilterPath(v ...string) func(*SecurityGetTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetToken) WithHeader(h map[string]string) func(*SecurityGetTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetToken) WithOpaqueID(s string) func(*SecurityGetTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}
