package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityQueryRoleFunc(t Transport) SecurityQueryRole {
	_ = "STUB: not implemented"
	return *new(SecurityQueryRole)
}

type SecurityQueryRole func(o ...func(*SecurityQueryRoleRequest)) (*Response, error)

type SecurityQueryRoleRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityQueryRoleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityQueryRole) WithContext(v context.Context) func(*SecurityQueryRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryRole) WithBody(v io.Reader) func(*SecurityQueryRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryRole) WithPretty() func(*SecurityQueryRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryRole) WithHuman() func(*SecurityQueryRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryRole) WithErrorTrace() func(*SecurityQueryRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryRole) WithFilterPath(v ...string) func(*SecurityQueryRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryRole) WithHeader(h map[string]string) func(*SecurityQueryRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryRole) WithOpaqueID(s string) func(*SecurityQueryRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}
