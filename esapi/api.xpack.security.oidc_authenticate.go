package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityOidcAuthenticateFunc(t Transport) SecurityOidcAuthenticate {
	_ = "STUB: not implemented"
	return *new(SecurityOidcAuthenticate)
}

type SecurityOidcAuthenticate func(body io.Reader, o ...func(*SecurityOidcAuthenticateRequest)) (*Response, error)

type SecurityOidcAuthenticateRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityOidcAuthenticateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityOidcAuthenticate) WithContext(v context.Context) func(*SecurityOidcAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcAuthenticate) WithPretty() func(*SecurityOidcAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcAuthenticate) WithHuman() func(*SecurityOidcAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcAuthenticate) WithErrorTrace() func(*SecurityOidcAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcAuthenticate) WithFilterPath(v ...string) func(*SecurityOidcAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcAuthenticate) WithHeader(h map[string]string) func(*SecurityOidcAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcAuthenticate) WithOpaqueID(s string) func(*SecurityOidcAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}
