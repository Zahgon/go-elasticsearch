package esapi

import (
	"context"
	"net/http"
)

func newSecurityDeleteRoleMappingFunc(t Transport) SecurityDeleteRoleMapping {
	_ = "STUB: not implemented"
	return *new(SecurityDeleteRoleMapping)
}

type SecurityDeleteRoleMapping func(name string, o ...func(*SecurityDeleteRoleMappingRequest)) (*Response, error)

type SecurityDeleteRoleMappingRequest struct {
	Name string

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityDeleteRoleMappingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityDeleteRoleMapping) WithContext(v context.Context) func(*SecurityDeleteRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRoleMapping) WithRefresh(v string) func(*SecurityDeleteRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRoleMapping) WithPretty() func(*SecurityDeleteRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRoleMapping) WithHuman() func(*SecurityDeleteRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRoleMapping) WithErrorTrace() func(*SecurityDeleteRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRoleMapping) WithFilterPath(v ...string) func(*SecurityDeleteRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRoleMapping) WithHeader(h map[string]string) func(*SecurityDeleteRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRoleMapping) WithOpaqueID(s string) func(*SecurityDeleteRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}
