package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityPutPrivilegesFunc(t Transport) SecurityPutPrivileges {
	_ = "STUB: not implemented"
	return *new(SecurityPutPrivileges)
}

type SecurityPutPrivileges func(body io.Reader, o ...func(*SecurityPutPrivilegesRequest)) (*Response, error)

type SecurityPutPrivilegesRequest struct {
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

func (r SecurityPutPrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityPutPrivileges) WithContext(v context.Context) func(*SecurityPutPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutPrivileges) WithRefresh(v string) func(*SecurityPutPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutPrivileges) WithPretty() func(*SecurityPutPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutPrivileges) WithHuman() func(*SecurityPutPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutPrivileges) WithErrorTrace() func(*SecurityPutPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutPrivileges) WithFilterPath(v ...string) func(*SecurityPutPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutPrivileges) WithHeader(h map[string]string) func(*SecurityPutPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityPutPrivileges) WithOpaqueID(s string) func(*SecurityPutPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}
