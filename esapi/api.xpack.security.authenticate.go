package esapi

import (
	"context"
	"net/http"
)

func newSecurityAuthenticateFunc(t Transport) SecurityAuthenticate {
	_ = "STUB: not implemented"
	return *new(SecurityAuthenticate)
}

type SecurityAuthenticate func(o ...func(*SecurityAuthenticateRequest)) (*Response, error)

type SecurityAuthenticateRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityAuthenticateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityAuthenticate) WithContext(v context.Context) func(*SecurityAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityAuthenticate) WithPretty() func(*SecurityAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityAuthenticate) WithHuman() func(*SecurityAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityAuthenticate) WithErrorTrace() func(*SecurityAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityAuthenticate) WithFilterPath(v ...string) func(*SecurityAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityAuthenticate) WithHeader(h map[string]string) func(*SecurityAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityAuthenticate) WithOpaqueID(s string) func(*SecurityAuthenticateRequest) {
	_ = "STUB: not implemented"
	return nil
}
