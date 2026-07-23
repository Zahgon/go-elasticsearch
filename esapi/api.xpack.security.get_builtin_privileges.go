package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetBuiltinPrivilegesFunc(t Transport) SecurityGetBuiltinPrivileges {
	_ = "STUB: not implemented"
	return *new(SecurityGetBuiltinPrivileges)
}

type SecurityGetBuiltinPrivileges func(o ...func(*SecurityGetBuiltinPrivilegesRequest)) (*Response, error)

type SecurityGetBuiltinPrivilegesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetBuiltinPrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetBuiltinPrivileges) WithContext(v context.Context) func(*SecurityGetBuiltinPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetBuiltinPrivileges) WithPretty() func(*SecurityGetBuiltinPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetBuiltinPrivileges) WithHuman() func(*SecurityGetBuiltinPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetBuiltinPrivileges) WithErrorTrace() func(*SecurityGetBuiltinPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetBuiltinPrivileges) WithFilterPath(v ...string) func(*SecurityGetBuiltinPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetBuiltinPrivileges) WithHeader(h map[string]string) func(*SecurityGetBuiltinPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetBuiltinPrivileges) WithOpaqueID(s string) func(*SecurityGetBuiltinPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}
