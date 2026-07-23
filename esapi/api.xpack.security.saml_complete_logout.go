package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecuritySamlCompleteLogoutFunc(t Transport) SecuritySamlCompleteLogout {
	_ = "STUB: not implemented"
	return *new(SecuritySamlCompleteLogout)
}

type SecuritySamlCompleteLogout func(body io.Reader, o ...func(*SecuritySamlCompleteLogoutRequest)) (*Response, error)

type SecuritySamlCompleteLogoutRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecuritySamlCompleteLogoutRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecuritySamlCompleteLogout) WithContext(v context.Context) func(*SecuritySamlCompleteLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlCompleteLogout) WithPretty() func(*SecuritySamlCompleteLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlCompleteLogout) WithHuman() func(*SecuritySamlCompleteLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlCompleteLogout) WithErrorTrace() func(*SecuritySamlCompleteLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlCompleteLogout) WithFilterPath(v ...string) func(*SecuritySamlCompleteLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlCompleteLogout) WithHeader(h map[string]string) func(*SecuritySamlCompleteLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlCompleteLogout) WithOpaqueID(s string) func(*SecuritySamlCompleteLogoutRequest) {
	_ = "STUB: not implemented"
	return nil
}
