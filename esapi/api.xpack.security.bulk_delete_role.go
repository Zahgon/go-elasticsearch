package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityBulkDeleteRoleFunc(t Transport) SecurityBulkDeleteRole {
	_ = "STUB: not implemented"
	return *new(SecurityBulkDeleteRole)
}

type SecurityBulkDeleteRole func(body io.Reader, o ...func(*SecurityBulkDeleteRoleRequest)) (*Response, error)

type SecurityBulkDeleteRoleRequest struct {
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

func (r SecurityBulkDeleteRoleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityBulkDeleteRole) WithContext(v context.Context) func(*SecurityBulkDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkDeleteRole) WithRefresh(v string) func(*SecurityBulkDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkDeleteRole) WithPretty() func(*SecurityBulkDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkDeleteRole) WithHuman() func(*SecurityBulkDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkDeleteRole) WithErrorTrace() func(*SecurityBulkDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkDeleteRole) WithFilterPath(v ...string) func(*SecurityBulkDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkDeleteRole) WithHeader(h map[string]string) func(*SecurityBulkDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkDeleteRole) WithOpaqueID(s string) func(*SecurityBulkDeleteRoleRequest) {
	_ = "STUB: not implemented"
	return nil
}
