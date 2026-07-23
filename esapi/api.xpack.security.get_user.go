package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetUserFunc(t Transport) SecurityGetUser {
	_ = "STUB: not implemented"
	return *new(SecurityGetUser)
}

type SecurityGetUser func(o ...func(*SecurityGetUserRequest)) (*Response, error)

type SecurityGetUserRequest struct {
	Username []string

	WithProfileUID *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetUserRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetUser) WithContext(v context.Context) func(*SecurityGetUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUser) WithUsername(v ...string) func(*SecurityGetUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUser) WithWithProfileUID(v bool) func(*SecurityGetUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUser) WithPretty() func(*SecurityGetUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUser) WithHuman() func(*SecurityGetUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUser) WithErrorTrace() func(*SecurityGetUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUser) WithFilterPath(v ...string) func(*SecurityGetUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUser) WithHeader(h map[string]string) func(*SecurityGetUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUser) WithOpaqueID(s string) func(*SecurityGetUserRequest) {
	_ = "STUB: not implemented"
	return nil
}
