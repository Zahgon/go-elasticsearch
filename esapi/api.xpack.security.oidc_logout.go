package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityOidcLogoutFunc(t Transport) SecurityOidcLogout {
	_ = "STUB: not implemented"
	return *new(SecurityOidcLogout)
}

type SecurityOidcLogout func(body io.Reader, o ...func(*SecurityOidcLogoutRequest)) (*Response, error)

type SecurityOidcLogoutRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityOidcLogoutRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityOidcLogout) WithContext(v context.Context) func(*SecurityOidcLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcLogout) WithPretty() func(*SecurityOidcLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcLogout) WithHuman() func(*SecurityOidcLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcLogout) WithErrorTrace() func(*SecurityOidcLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcLogout) WithFilterPath(v ...string) func(*SecurityOidcLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcLogout) WithHeader(h map[string]string) func(*SecurityOidcLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcLogout) WithOpaqueID(s string) func(*SecurityOidcLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}
