package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityCloneAPIKeyFunc(t Transport) SecurityCloneAPIKey {
	_ = "STUB: not implemented"
	return *new(SecurityCloneAPIKey)
}

type SecurityCloneAPIKey func(body io.Reader, o ...func(*SecurityCloneAPIKeyRequest)) (*Response, error)

type SecurityCloneAPIKeyRequest struct {
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

func (r SecurityCloneAPIKeyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityCloneAPIKey) WithContext(v context.Context) func(*SecurityCloneAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCloneAPIKey) WithRefresh(v string) func(*SecurityCloneAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCloneAPIKey) WithPretty() func(*SecurityCloneAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCloneAPIKey) WithHuman() func(*SecurityCloneAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCloneAPIKey) WithErrorTrace() func(*SecurityCloneAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCloneAPIKey) WithFilterPath(v ...string) func(*SecurityCloneAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCloneAPIKey) WithHeader(h map[string]string) func(*SecurityCloneAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCloneAPIKey) WithOpaqueID(s string) func(*SecurityCloneAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}
