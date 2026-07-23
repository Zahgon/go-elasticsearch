package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityOidcPrepareAuthenticationFunc(t Transport) SecurityOidcPrepareAuthentication {
	_ = "STUB: not implemented"
	return *new(SecurityOidcPrepareAuthentication)
}

type SecurityOidcPrepareAuthentication func(body io.Reader, o ...func(*SecurityOidcPrepareAuthenticationRequest)) (*Response, error)

type SecurityOidcPrepareAuthenticationRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityOidcPrepareAuthenticationRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityOidcPrepareAuthentication) WithContext(v context.Context) func(*SecurityOidcPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcPrepareAuthentication) WithPretty() func(*SecurityOidcPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcPrepareAuthentication) WithHuman() func(*SecurityOidcPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcPrepareAuthentication) WithErrorTrace() func(*SecurityOidcPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcPrepareAuthentication) WithFilterPath(v ...string) func(*SecurityOidcPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcPrepareAuthentication) WithHeader(h map[string]string) func(*SecurityOidcPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityOidcPrepareAuthentication) WithOpaqueID(s string) func(*SecurityOidcPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}
