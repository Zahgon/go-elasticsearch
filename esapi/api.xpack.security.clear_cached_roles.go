package esapi

import (
	"context"
	"net/http"
)

func newSecurityClearCachedRolesFunc(t Transport) SecurityClearCachedRoles {
	_ = "STUB: not implemented"
	return *new(SecurityClearCachedRoles)
}

type SecurityClearCachedRoles func(name []string, o ...func(*SecurityClearCachedRolesRequest)) (*Response, error)

type SecurityClearCachedRolesRequest struct {
	Name []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityClearCachedRolesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityClearCachedRoles) WithContext(v context.Context) func(*SecurityClearCachedRolesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRoles) WithPretty() func(*SecurityClearCachedRolesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRoles) WithHuman() func(*SecurityClearCachedRolesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRoles) WithErrorTrace() func(*SecurityClearCachedRolesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRoles) WithFilterPath(v ...string) func(*SecurityClearCachedRolesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRoles) WithHeader(h map[string]string) func(*SecurityClearCachedRolesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRoles) WithOpaqueID(s string) func(*SecurityClearCachedRolesRequest) {
	_ = "STUB: not implemented"
	return nil
}
