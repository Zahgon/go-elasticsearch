package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityBulkPutRoleFunc(t Transport) SecurityBulkPutRole {
	_ = "STUB: not implemented"
	return *new(SecurityBulkPutRole)
}

type SecurityBulkPutRole func(body io.Reader, o ...func(*SecurityBulkPutRoleRequest)) (*Response, error)

type SecurityBulkPutRoleRequest struct {
	Body io.Reader

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityBulkPutRoleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityBulkPutRole) WithContext(v context.Context) func(*SecurityBulkPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkPutRole) WithRefresh(v string) func(*SecurityBulkPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkPutRole) WithPretty() func(*SecurityBulkPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkPutRole) WithHuman() func(*SecurityBulkPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkPutRole) WithErrorTrace() func(*SecurityBulkPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkPutRole) WithFilterPath(v ...string) func(*SecurityBulkPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkPutRole) WithHeader(h map[string]string) func(*SecurityBulkPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkPutRole) WithOpaqueID(s string) func(*SecurityBulkPutRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}
