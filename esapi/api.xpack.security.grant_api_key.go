package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityGrantAPIKeyFunc(t Transport) SecurityGrantAPIKey {
	_ = "STUB: not implemented"
	return *new(SecurityGrantAPIKey)
}

type SecurityGrantAPIKey func(body io.Reader, o ...func(*SecurityGrantAPIKeyRequest)) (*Response, error)

type SecurityGrantAPIKeyRequest struct {
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

func (r SecurityGrantAPIKeyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGrantAPIKey) WithContext(v context.Context) func(*SecurityGrantAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGrantAPIKey) WithRefresh(v string) func(*SecurityGrantAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGrantAPIKey) WithPretty() func(*SecurityGrantAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGrantAPIKey) WithHuman() func(*SecurityGrantAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGrantAPIKey) WithErrorTrace() func(*SecurityGrantAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGrantAPIKey) WithFilterPath(v ...string) func(*SecurityGrantAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGrantAPIKey) WithHeader(h map[string]string) func(*SecurityGrantAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGrantAPIKey) WithOpaqueID(s string) func(*SecurityGrantAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}
