package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityInvalidateAPIKeyFunc(t Transport) SecurityInvalidateAPIKey {
	_ = "STUB: not implemented"
	return *new(SecurityInvalidateAPIKey)
}

type SecurityInvalidateAPIKey func(body io.Reader, o ...func(*SecurityInvalidateAPIKeyRequest)) (*Response, error)

type SecurityInvalidateAPIKeyRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityInvalidateAPIKeyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityInvalidateAPIKey) WithContext(v context.Context) func(*SecurityInvalidateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateAPIKey) WithPretty() func(*SecurityInvalidateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateAPIKey) WithHuman() func(*SecurityInvalidateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateAPIKey) WithErrorTrace() func(*SecurityInvalidateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateAPIKey) WithFilterPath(v ...string) func(*SecurityInvalidateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateAPIKey) WithHeader(h map[string]string) func(*SecurityInvalidateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityInvalidateAPIKey) WithOpaqueID(s string) func(*SecurityInvalidateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}
