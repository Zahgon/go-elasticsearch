package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetServiceCredentialsFunc(t Transport) SecurityGetServiceCredentials {
	_ = "STUB: not implemented"
	return *new(SecurityGetServiceCredentials)
}

type SecurityGetServiceCredentials func(namespace string, service string, o ...func(*SecurityGetServiceCredentialsRequest)) (*Response, error)

type SecurityGetServiceCredentialsRequest struct {
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

func (r SecurityGetServiceCredentialsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetServiceCredentials) WithContext(v context.Context) func(*SecurityGetServiceCredentialsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceCredentials) WithPretty() func(*SecurityGetServiceCredentialsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceCredentials) WithHuman() func(*SecurityGetServiceCredentialsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceCredentials) WithErrorTrace() func(*SecurityGetServiceCredentialsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceCredentials) WithFilterPath(v ...string) func(*SecurityGetServiceCredentialsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceCredentials) WithHeader(h map[string]string) func(*SecurityGetServiceCredentialsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetServiceCredentials) WithOpaqueID(s string) func(*SecurityGetServiceCredentialsRequest) {
	_ = "STUB: not implemented"
	return nil
}
