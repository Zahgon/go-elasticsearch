package esapi

import (
	"context"
	"net/http"
)

func newSecurityDeleteRoleFunc(t Transport) SecurityDeleteRole {
	_ = "STUB: not implemented"
	return *new(SecurityDeleteRole)
}

type SecurityDeleteRole func(name string, o ...func(*SecurityDeleteRoleRequest)) (*Response, error)

type SecurityDeleteRoleRequest struct {
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

func (r SecurityDeleteRoleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityDeleteRole) WithContext(v context.Context) func(*SecurityDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRole) WithRefresh(v string) func(*SecurityDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRole) WithPretty() func(*SecurityDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRole) WithHuman() func(*SecurityDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRole) WithErrorTrace() func(*SecurityDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRole) WithFilterPath(v ...string) func(*SecurityDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRole) WithHeader(h map[string]string) func(*SecurityDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteRole) WithOpaqueID(s string) func(*SecurityDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}
