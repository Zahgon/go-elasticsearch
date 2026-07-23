package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityQueryUserFunc(t Transport) SecurityQueryUser {
	_ = "STUB: not implemented"
	return *new(SecurityQueryUser)
}

type SecurityQueryUser func(o ...func(*SecurityQueryUserRequest)) (*Response, error)

type SecurityQueryUserRequest struct {
	Body io.Reader

	WithProfileUID *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityQueryUserRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityQueryUser) WithContext(v context.Context) func(*SecurityQueryUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryUser) WithBody(v io.Reader) func(*SecurityQueryUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryUser) WithWithProfileUID(v bool) func(*SecurityQueryUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryUser) WithPretty() func(*SecurityQueryUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryUser) WithHuman() func(*SecurityQueryUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryUser) WithErrorTrace() func(*SecurityQueryUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryUser) WithFilterPath(v ...string) func(*SecurityQueryUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryUser) WithHeader(h map[string]string) func(*SecurityQueryUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryUser) WithOpaqueID(s string) func(*SecurityQueryUserRequest) {
	_ = "STUB: not implemented"
	return nil
}
