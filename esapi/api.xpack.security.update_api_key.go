package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityUpdateAPIKeyFunc(t Transport) SecurityUpdateAPIKey {
	_ = "STUB: not implemented"
	return *new(SecurityUpdateAPIKey)
}

type SecurityUpdateAPIKey func(id string, o ...func(*SecurityUpdateAPIKeyRequest)) (*Response, error)

type SecurityUpdateAPIKeyRequest struct {
	DocumentID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityUpdateAPIKeyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityUpdateAPIKey) WithContext(v context.Context) func(*SecurityUpdateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateAPIKey) WithBody(v io.Reader) func(*SecurityUpdateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateAPIKey) WithPretty() func(*SecurityUpdateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateAPIKey) WithHuman() func(*SecurityUpdateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateAPIKey) WithErrorTrace() func(*SecurityUpdateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateAPIKey) WithFilterPath(v ...string) func(*SecurityUpdateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateAPIKey) WithHeader(h map[string]string) func(*SecurityUpdateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateAPIKey) WithOpaqueID(s string) func(*SecurityUpdateAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}
