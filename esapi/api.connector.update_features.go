package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateFeaturesFunc(t Transport) ConnectorUpdateFeatures {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateFeatures)
}

type ConnectorUpdateFeatures func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateFeaturesRequest)) (*Response, error)

type ConnectorUpdateFeaturesRequest struct {
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

func (r ConnectorUpdateFeaturesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateFeatures) WithContext(v context.Context) func(*ConnectorUpdateFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFeatures) WithPretty() func(*ConnectorUpdateFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFeatures) WithHuman() func(*ConnectorUpdateFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFeatures) WithErrorTrace() func(*ConnectorUpdateFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFeatures) WithFilterPath(v ...string) func(*ConnectorUpdateFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFeatures) WithHeader(h map[string]string) func(*ConnectorUpdateFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFeatures) WithOpaqueID(s string) func(*ConnectorUpdateFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}
