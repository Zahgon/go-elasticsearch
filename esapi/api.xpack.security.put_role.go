package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityPutRoleFunc(t Transport) SecurityPutRole {
	_ = "STUB: not implemented"
	return *new(SecurityPutRole)
}

type SecurityPutRole func(name string, body io.Reader, o ...func(*SecurityPutRoleRequest)) (*Response, error)

type SecurityPutRoleRequest struct {
	Body io.Reader

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

func (r SecurityPutRoleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityPutRole) WithContext(v context.Context) func(*SecurityPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRole) WithRefresh(v string) func(*SecurityPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRole) WithPretty() func(*SecurityPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRole) WithHuman() func(*SecurityPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRole) WithErrorTrace() func(*SecurityPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRole) WithFilterPath(v ...string) func(*SecurityPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRole) WithHeader(h map[string]string) func(*SecurityPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRole) WithOpaqueID(s string) func(*SecurityPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}
