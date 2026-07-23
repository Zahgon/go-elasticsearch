package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetRoleFunc(t Transport) SecurityGetRole {
	_ = "STUB: not implemented"
	return *new(SecurityGetRole)
}

type SecurityGetRole func(o ...func(*SecurityGetRoleRequest)) (*Response, error)

type SecurityGetRoleRequest struct {
	Name []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetRoleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetRole) WithContext(v context.Context) func(*SecurityGetRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRole) WithName(v ...string) func(*SecurityGetRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRole) WithPretty() func(*SecurityGetRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRole) WithHuman() func(*SecurityGetRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRole) WithErrorTrace() func(*SecurityGetRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRole) WithFilterPath(v ...string) func(*SecurityGetRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRole) WithHeader(h map[string]string) func(*SecurityGetRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetRole) WithOpaqueID(s string) func(*SecurityGetRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}
