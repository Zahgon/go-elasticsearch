package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityPutRoleMappingFunc(t Transport) SecurityPutRoleMapping {
	_ = "STUB: not implemented"
	return *new(SecurityPutRoleMapping)
}

type SecurityPutRoleMapping func(name string, body io.Reader, o ...func(*SecurityPutRoleMappingRequest)) (*Response, error)

type SecurityPutRoleMappingRequest struct {
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

func (r SecurityPutRoleMappingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityPutRoleMapping) WithContext(v context.Context) func(*SecurityPutRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRoleMapping) WithRefresh(v string) func(*SecurityPutRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRoleMapping) WithPretty() func(*SecurityPutRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRoleMapping) WithHuman() func(*SecurityPutRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRoleMapping) WithErrorTrace() func(*SecurityPutRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRoleMapping) WithFilterPath(v ...string) func(*SecurityPutRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRoleMapping) WithHeader(h map[string]string) func(*SecurityPutRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutRoleMapping) WithOpaqueID(s string) func(*SecurityPutRoleMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}
