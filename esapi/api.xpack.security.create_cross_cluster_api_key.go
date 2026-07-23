package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityCreateCrossClusterAPIKeyFunc(t Transport) SecurityCreateCrossClusterAPIKey {
	_ = "STUB: not implemented"
	return *new(SecurityCreateCrossClusterAPIKey)
}

type SecurityCreateCrossClusterAPIKey func(body io.Reader, o ...func(*SecurityCreateCrossClusterAPIKeyRequest)) (*Response, error)

type SecurityCreateCrossClusterAPIKeyRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityCreateCrossClusterAPIKeyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityCreateCrossClusterAPIKey) WithContext(v context.Context) func(*SecurityCreateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateCrossClusterAPIKey) WithPretty() func(*SecurityCreateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateCrossClusterAPIKey) WithHuman() func(*SecurityCreateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateCrossClusterAPIKey) WithErrorTrace() func(*SecurityCreateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateCrossClusterAPIKey) WithFilterPath(v ...string) func(*SecurityCreateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateCrossClusterAPIKey) WithHeader(h map[string]string) func(*SecurityCreateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityCreateCrossClusterAPIKey) WithOpaqueID(s string) func(*SecurityCreateCrossClusterAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}
