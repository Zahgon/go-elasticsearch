package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityCreateAPIKeyFunc(t Transport) SecurityCreateAPIKey {
	_ = "STUB: not implemented"
	return *new(SecurityCreateAPIKey)
}

type SecurityCreateAPIKey func(body io.Reader, o ...func(*SecurityCreateAPIKeyRequest)) (*Response, error)

type SecurityCreateAPIKeyRequest struct {
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

func (r SecurityCreateAPIKeyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityCreateAPIKey) WithContext(v context.Context) func(*SecurityCreateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateAPIKey) WithRefresh(v string) func(*SecurityCreateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateAPIKey) WithPretty() func(*SecurityCreateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateAPIKey) WithHuman() func(*SecurityCreateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateAPIKey) WithErrorTrace() func(*SecurityCreateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateAPIKey) WithFilterPath(v ...string) func(*SecurityCreateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateAPIKey) WithHeader(h map[string]string) func(*SecurityCreateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateAPIKey) WithOpaqueID(s string) func(*SecurityCreateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}
