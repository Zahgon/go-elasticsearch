package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityUpdateCrossClusterAPIKeyFunc(t Transport) SecurityUpdateCrossClusterAPIKey {
	_ = "STUB: not implemented"
	return *new(SecurityUpdateCrossClusterAPIKey)
}

type SecurityUpdateCrossClusterAPIKey func(id string, body io.Reader, o ...func(*SecurityUpdateCrossClusterAPIKeyRequest)) (*Response, error)

type SecurityUpdateCrossClusterAPIKeyRequest struct {
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

func (r SecurityUpdateCrossClusterAPIKeyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityUpdateCrossClusterAPIKey) WithContext(v context.Context) func(*SecurityUpdateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateCrossClusterAPIKey) WithPretty() func(*SecurityUpdateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateCrossClusterAPIKey) WithHuman() func(*SecurityUpdateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateCrossClusterAPIKey) WithErrorTrace() func(*SecurityUpdateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateCrossClusterAPIKey) WithFilterPath(v ...string) func(*SecurityUpdateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateCrossClusterAPIKey) WithHeader(h map[string]string) func(*SecurityUpdateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateCrossClusterAPIKey) WithOpaqueID(s string) func(*SecurityUpdateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}
