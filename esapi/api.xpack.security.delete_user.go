package esapi

import (
	"context"
	"net/http"
)

func newSecurityDeleteUserFunc(t Transport) SecurityDeleteUser {
	_ = "STUB: not implemented"
	return *new(SecurityDeleteUser)
}

type SecurityDeleteUser func(username string, o ...func(*SecurityDeleteUserRequest)) (*Response, error)

type SecurityDeleteUserRequest struct {
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

func (r SecurityDeleteUserRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityDeleteUser) WithContext(v context.Context) func(*SecurityDeleteUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteUser) WithRefresh(v string) func(*SecurityDeleteUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteUser) WithPretty() func(*SecurityDeleteUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteUser) WithHuman() func(*SecurityDeleteUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteUser) WithErrorTrace() func(*SecurityDeleteUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteUser) WithFilterPath(v ...string) func(*SecurityDeleteUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteUser) WithHeader(h map[string]string) func(*SecurityDeleteUserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteUser) WithOpaqueID(s string) func(*SecurityDeleteUserRequest) {
	_ = "STUB: not implemented"
	return nil
}
