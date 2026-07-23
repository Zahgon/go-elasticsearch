package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetUserPrivilegesFunc(t Transport) SecurityGetUserPrivileges {
	_ = "STUB: not implemented"
	return *new(SecurityGetUserPrivileges)
}

type SecurityGetUserPrivileges func(o ...func(*SecurityGetUserPrivilegesRequest)) (*Response, error)

type SecurityGetUserPrivilegesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetUserPrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetUserPrivileges) WithContext(v context.Context) func(*SecurityGetUserPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserPrivileges) WithPretty() func(*SecurityGetUserPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserPrivileges) WithHuman() func(*SecurityGetUserPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserPrivileges) WithErrorTrace() func(*SecurityGetUserPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserPrivileges) WithFilterPath(v ...string) func(*SecurityGetUserPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserPrivileges) WithHeader(h map[string]string) func(*SecurityGetUserPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserPrivileges) WithOpaqueID(s string) func(*SecurityGetUserPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}
