package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityHasPrivilegesFunc(t Transport) SecurityHasPrivileges {
	_ = "STUB: not implemented"
	return *new(SecurityHasPrivileges)
}

type SecurityHasPrivileges func(body io.Reader, o ...func(*SecurityHasPrivilegesRequest)) (*Response, error)

type SecurityHasPrivilegesRequest struct {
	Body io.Reader

	User string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityHasPrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityHasPrivileges) WithContext(v context.Context) func(*SecurityHasPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivileges) WithUser(v string) func(*SecurityHasPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivileges) WithPretty() func(*SecurityHasPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivileges) WithHuman() func(*SecurityHasPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivileges) WithErrorTrace() func(*SecurityHasPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivileges) WithFilterPath(v ...string) func(*SecurityHasPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivileges) WithHeader(h map[string]string) func(*SecurityHasPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivileges) WithOpaqueID(s string) func(*SecurityHasPrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}
