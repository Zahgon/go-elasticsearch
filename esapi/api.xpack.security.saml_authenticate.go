package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecuritySamlAuthenticateFunc(t Transport) SecuritySamlAuthenticate {
	_ = "STUB: not implemented"
	return *new(SecuritySamlAuthenticate)
}

type SecuritySamlAuthenticate func(body io.Reader, o ...func(*SecuritySamlAuthenticateRequest)) (*Response, error)

type SecuritySamlAuthenticateRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecuritySamlAuthenticateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecuritySamlAuthenticate) WithContext(v context.Context) func(*SecuritySamlAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlAuthenticate) WithPretty() func(*SecuritySamlAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlAuthenticate) WithHuman() func(*SecuritySamlAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlAuthenticate) WithErrorTrace() func(*SecuritySamlAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlAuthenticate) WithFilterPath(v ...string) func(*SecuritySamlAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlAuthenticate) WithHeader(h map[string]string) func(*SecuritySamlAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlAuthenticate) WithOpaqueID(s string) func(*SecuritySamlAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}
