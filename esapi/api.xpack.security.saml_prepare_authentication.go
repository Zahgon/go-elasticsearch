package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecuritySamlPrepareAuthenticationFunc(t Transport) SecuritySamlPrepareAuthentication {
	_ = "STUB: not implemented"
	return *new(SecuritySamlPrepareAuthentication)
}

type SecuritySamlPrepareAuthentication func(body io.Reader, o ...func(*SecuritySamlPrepareAuthenticationRequest)) (*Response, error)

type SecuritySamlPrepareAuthenticationRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecuritySamlPrepareAuthenticationRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecuritySamlPrepareAuthentication) WithContext(v context.Context) func(*SecuritySamlPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlPrepareAuthentication) WithPretty() func(*SecuritySamlPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlPrepareAuthentication) WithHuman() func(*SecuritySamlPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlPrepareAuthentication) WithErrorTrace() func(*SecuritySamlPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlPrepareAuthentication) WithFilterPath(v ...string) func(*SecuritySamlPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlPrepareAuthentication) WithHeader(h map[string]string) func(*SecuritySamlPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlPrepareAuthentication) WithOpaqueID(s string) func(*SecuritySamlPrepareAuthenticationRequest) {
	_ = "STUB: not implemented"
	return nil
}
