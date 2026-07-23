package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetServiceAccountsFunc(t Transport) SecurityGetServiceAccounts {
	_ = "STUB: not implemented"
	return *new(SecurityGetServiceAccounts)
}

type SecurityGetServiceAccounts func(o ...func(*SecurityGetServiceAccountsRequest)) (*Response, error)

type SecurityGetServiceAccountsRequest struct {
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

func (r SecurityGetServiceAccountsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetServiceAccounts) WithContext(v context.Context) func(*SecurityGetServiceAccountsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceAccounts) WithNamespace(v string) func(*SecurityGetServiceAccountsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceAccounts) WithService(v string) func(*SecurityGetServiceAccountsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceAccounts) WithPretty() func(*SecurityGetServiceAccountsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceAccounts) WithHuman() func(*SecurityGetServiceAccountsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceAccounts) WithErrorTrace() func(*SecurityGetServiceAccountsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceAccounts) WithFilterPath(v ...string) func(*SecurityGetServiceAccountsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceAccounts) WithHeader(h map[string]string) func(*SecurityGetServiceAccountsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceAccounts) WithOpaqueID(s string) func(*SecurityGetServiceAccountsRequest) {
	_ = "STUB: not implemented"
	return nil
}
