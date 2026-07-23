package esapi

import (
	"context"
	"net/http"
)

func newSecurityDeleteServiceTokenFunc(t Transport) SecurityDeleteServiceToken {
	_ = "STUB: not implemented"
	return *new(SecurityDeleteServiceToken)
}

type SecurityDeleteServiceToken func(name string, namespace string, service string, o ...func(*SecurityDeleteServiceTokenRequest)) (*Response, error)

type SecurityDeleteServiceTokenRequest struct {
	Name      string
	Namespace string
	Service   string

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityDeleteServiceTokenRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityDeleteServiceToken) WithContext(v context.Context) func(*SecurityDeleteServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteServiceToken) WithRefresh(v string) func(*SecurityDeleteServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteServiceToken) WithPretty() func(*SecurityDeleteServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteServiceToken) WithHuman() func(*SecurityDeleteServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteServiceToken) WithErrorTrace() func(*SecurityDeleteServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteServiceToken) WithFilterPath(v ...string) func(*SecurityDeleteServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteServiceToken) WithHeader(h map[string]string) func(*SecurityDeleteServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDeleteServiceToken) WithOpaqueID(s string) func(*SecurityDeleteServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}
