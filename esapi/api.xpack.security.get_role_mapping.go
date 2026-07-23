package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetRoleMappingFunc(t Transport) SecurityGetRoleMapping {
	_ = "STUB: not implemented"
	return *new(SecurityGetRoleMapping)
}

type SecurityGetRoleMapping func(o ...func(*SecurityGetRoleMappingRequest)) (*Response, error)

type SecurityGetRoleMappingRequest struct {
	Name []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetRoleMappingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetRoleMapping) WithContext(v context.Context) func(*SecurityGetRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRoleMapping) WithName(v ...string) func(*SecurityGetRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRoleMapping) WithPretty() func(*SecurityGetRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRoleMapping) WithHuman() func(*SecurityGetRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRoleMapping) WithErrorTrace() func(*SecurityGetRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRoleMapping) WithFilterPath(v ...string) func(*SecurityGetRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRoleMapping) WithHeader(h map[string]string) func(*SecurityGetRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRoleMapping) WithOpaqueID(s string) func(*SecurityGetRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}
