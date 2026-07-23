package esapi

import (
	"context"
	"net/http"
)

func newSecurityDeletePrivilegesFunc(t Transport) SecurityDeletePrivileges {
	_ = "STUB: not implemented"
	return *new(SecurityDeletePrivileges)
}

type SecurityDeletePrivileges func(name []string, application string, o ...func(*SecurityDeletePrivilegesRequest)) (*Response, error)

type SecurityDeletePrivilegesRequest struct {
	Application string
	Name        []string

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityDeletePrivilegesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityDeletePrivileges) WithContext(v context.Context) func(*SecurityDeletePrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeletePrivileges) WithRefresh(v string) func(*SecurityDeletePrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeletePrivileges) WithPretty() func(*SecurityDeletePrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeletePrivileges) WithHuman() func(*SecurityDeletePrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeletePrivileges) WithErrorTrace() func(*SecurityDeletePrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeletePrivileges) WithFilterPath(v ...string) func(*SecurityDeletePrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeletePrivileges) WithHeader(h map[string]string) func(*SecurityDeletePrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeletePrivileges) WithOpaqueID(s string) func(*SecurityDeletePrivilegesRequest) {
	_ = "STUB: not implemented"
	return nil
}
