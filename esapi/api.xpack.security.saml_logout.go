package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecuritySamlLogoutFunc(t Transport) SecuritySamlLogout {
	_ = "STUB: not implemented"
	return *new(SecuritySamlLogout)
}

type SecuritySamlLogout func(body io.Reader, o ...func(*SecuritySamlLogoutRequest)) (*Response, error)

type SecuritySamlLogoutRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecuritySamlLogoutRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecuritySamlLogout) WithContext(v context.Context) func(*SecuritySamlLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlLogout) WithPretty() func(*SecuritySamlLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlLogout) WithHuman() func(*SecuritySamlLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlLogout) WithErrorTrace() func(*SecuritySamlLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlLogout) WithFilterPath(v ...string) func(*SecuritySamlLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlLogout) WithHeader(h map[string]string) func(*SecuritySamlLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlLogout) WithOpaqueID(s string) func(*SecuritySamlLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}
