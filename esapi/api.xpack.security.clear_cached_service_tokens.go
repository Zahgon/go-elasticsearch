package esapi

import (
	"context"
	"net/http"
)

func newSecurityClearCachedServiceTokensFunc(t Transport) SecurityClearCachedServiceTokens {
	_ = "STUB: not implemented"
	return *new(SecurityClearCachedServiceTokens)
}

type SecurityClearCachedServiceTokens func(name []string, namespace string, service string, o ...func(*SecurityClearCachedServiceTokensRequest)) (*Response, error)

type SecurityClearCachedServiceTokensRequest struct {
	Name      []string
	Namespace string
	Service   string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityClearCachedServiceTokensRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityClearCachedServiceTokens) WithContext(v context.Context) func(*SecurityClearCachedServiceTokensRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedServiceTokens) WithPretty() func(*SecurityClearCachedServiceTokensRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedServiceTokens) WithHuman() func(*SecurityClearCachedServiceTokensRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedServiceTokens) WithErrorTrace() func(*SecurityClearCachedServiceTokensRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedServiceTokens) WithFilterPath(v ...string) func(*SecurityClearCachedServiceTokensRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedServiceTokens) WithHeader(h map[string]string) func(*SecurityClearCachedServiceTokensRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedServiceTokens) WithOpaqueID(s string) func(*SecurityClearCachedServiceTokensRequest) {
	_ = "STUB: not implemented"
	return nil
}
