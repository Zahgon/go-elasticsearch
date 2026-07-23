package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityChangePasswordFunc(t Transport) SecurityChangePassword {
	_ = "STUB: not implemented"
	return *new(SecurityChangePassword)
}

type SecurityChangePassword func(body io.Reader, o ...func(*SecurityChangePasswordRequest)) (*Response, error)

type SecurityChangePasswordRequest struct {
	Body io.Reader

	Username string

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityChangePasswordRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityChangePassword) WithContext(v context.Context) func(*SecurityChangePasswordRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityChangePassword) WithUsername(v string) func(*SecurityChangePasswordRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityChangePassword) WithRefresh(v string) func(*SecurityChangePasswordRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityChangePassword) WithPretty() func(*SecurityChangePasswordRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityChangePassword) WithHuman() func(*SecurityChangePasswordRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityChangePassword) WithErrorTrace() func(*SecurityChangePasswordRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityChangePassword) WithFilterPath(v ...string) func(*SecurityChangePasswordRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityChangePassword) WithHeader(h map[string]string) func(*SecurityChangePasswordRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityChangePassword) WithOpaqueID(s string) func(*SecurityChangePasswordRequest) {
	_ = "STUB: not implemented"
	return nil
}
