package esapi

import (
	"context"
	"net/http"
)

func newSecurityDisableUserFunc(t Transport) SecurityDisableUser {
	_ = "STUB: not implemented"
	return *new(SecurityDisableUser)
}

type SecurityDisableUser func(username string, o ...func(*SecurityDisableUserRequest)) (*Response, error)

type SecurityDisableUserRequest struct {
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

func (r SecurityDisableUserRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityDisableUser) WithContext(v context.Context) func(*SecurityDisableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUser) WithRefresh(v string) func(*SecurityDisableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUser) WithPretty() func(*SecurityDisableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUser) WithHuman() func(*SecurityDisableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUser) WithErrorTrace() func(*SecurityDisableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUser) WithFilterPath(v ...string) func(*SecurityDisableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUser) WithHeader(h map[string]string) func(*SecurityDisableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUser) WithOpaqueID(s string) func(*SecurityDisableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}
