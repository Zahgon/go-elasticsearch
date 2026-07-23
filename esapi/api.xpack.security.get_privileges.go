package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetPrivilegesFunc(t Transport) SecurityGetPrivileges {
	_ = "STUB: not implemented"
	return *new(SecurityGetPrivileges)
}

type SecurityGetPrivileges func(o ...func(*SecurityGetPrivilegesRequest)) (*Response, error)

type SecurityGetPrivilegesRequest struct {
	Application string
	Name        []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetPrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetPrivileges) WithContext(v context.Context) func(*SecurityGetPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetPrivileges) WithApplication(v string) func(*SecurityGetPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetPrivileges) WithName(v ...string) func(*SecurityGetPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetPrivileges) WithPretty() func(*SecurityGetPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetPrivileges) WithHuman() func(*SecurityGetPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetPrivileges) WithErrorTrace() func(*SecurityGetPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetPrivileges) WithFilterPath(v ...string) func(*SecurityGetPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetPrivileges) WithHeader(h map[string]string) func(*SecurityGetPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetPrivileges) WithOpaqueID(s string) func(*SecurityGetPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}
