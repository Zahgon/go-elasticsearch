package esapi

import (
	"context"
	"net/http"
)

func newSecurityCreateServiceTokenFunc(t Transport) SecurityCreateServiceToken {
	_ = "STUB: not implemented"
	return *new(SecurityCreateServiceToken)
}

type SecurityCreateServiceToken func(namespace string, service string, o ...func(*SecurityCreateServiceTokenRequest)) (*Response, error)

type SecurityCreateServiceTokenRequest struct {
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

func (r SecurityCreateServiceTokenRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityCreateServiceToken) WithContext(v context.Context) func(*SecurityCreateServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateServiceToken) WithName(v string) func(*SecurityCreateServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateServiceToken) WithRefresh(v string) func(*SecurityCreateServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateServiceToken) WithPretty() func(*SecurityCreateServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateServiceToken) WithHuman() func(*SecurityCreateServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateServiceToken) WithErrorTrace() func(*SecurityCreateServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateServiceToken) WithFilterPath(v ...string) func(*SecurityCreateServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateServiceToken) WithHeader(h map[string]string) func(*SecurityCreateServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateServiceToken) WithOpaqueID(s string) func(*SecurityCreateServiceTokenRequest) {
	_ = "STUB: not implemented"
	return nil
}
