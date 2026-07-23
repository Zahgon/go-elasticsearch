package esapi

import (
	"context"
	"net/http"
)

func newSecurityEnableUserFunc(t Transport) SecurityEnableUser {
	_ = "STUB: not implemented"
	return *new(SecurityEnableUser)
}

type SecurityEnableUser func(username string, o ...func(*SecurityEnableUserRequest)) (*Response, error)

type SecurityEnableUserRequest struct {
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

func (r SecurityEnableUserRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityEnableUser) WithContext(v context.Context) func(*SecurityEnableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUser) WithRefresh(v string) func(*SecurityEnableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUser) WithPretty() func(*SecurityEnableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUser) WithHuman() func(*SecurityEnableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUser) WithErrorTrace() func(*SecurityEnableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUser) WithFilterPath(v ...string) func(*SecurityEnableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUser) WithHeader(h map[string]string) func(*SecurityEnableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUser) WithOpaqueID(s string) func(*SecurityEnableUserRequest) {
	_ = "STUB: not implemented"
	return nil
}
