package esapi

import (
	"context"
	"net/http"
)

func newSecurityClearCachedPrivilegesFunc(t Transport) SecurityClearCachedPrivileges {
	_ = "STUB: not implemented"
	return *new(SecurityClearCachedPrivileges)
}

type SecurityClearCachedPrivileges func(application []string, o ...func(*SecurityClearCachedPrivilegesRequest)) (*Response, error)

type SecurityClearCachedPrivilegesRequest struct {
	Application []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityClearCachedPrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityClearCachedPrivileges) WithContext(v context.Context) func(*SecurityClearCachedPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedPrivileges) WithPretty() func(*SecurityClearCachedPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedPrivileges) WithHuman() func(*SecurityClearCachedPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedPrivileges) WithErrorTrace() func(*SecurityClearCachedPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedPrivileges) WithFilterPath(v ...string) func(*SecurityClearCachedPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedPrivileges) WithHeader(h map[string]string) func(*SecurityClearCachedPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedPrivileges) WithOpaqueID(s string) func(*SecurityClearCachedPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}
