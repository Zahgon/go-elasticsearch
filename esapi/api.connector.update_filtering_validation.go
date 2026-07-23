package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateFilteringValidationFunc(t Transport) ConnectorUpdateFilteringValidation {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateFilteringValidation)
}

type ConnectorUpdateFilteringValidation func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateFilteringValidationRequest)) (*Response, error)

type ConnectorUpdateFilteringValidationRequest struct {
	Body io.Reader

	ConnectorID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorUpdateFilteringValidationRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateFilteringValidation) WithContext(v context.Context) func(*ConnectorUpdateFilteringValidationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFilteringValidation) WithPretty() func(*ConnectorUpdateFilteringValidationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFilteringValidation) WithHuman() func(*ConnectorUpdateFilteringValidationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFilteringValidation) WithErrorTrace() func(*ConnectorUpdateFilteringValidationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFilteringValidation) WithFilterPath(v ...string) func(*ConnectorUpdateFilteringValidationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFilteringValidation) WithHeader(h map[string]string) func(*ConnectorUpdateFilteringValidationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFilteringValidation) WithOpaqueID(s string) func(*ConnectorUpdateFilteringValidationRequest) {
	_ = "STUB: not implemented"
	return nil
}
