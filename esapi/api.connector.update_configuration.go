package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateConfigurationFunc(t Transport) ConnectorUpdateConfiguration {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateConfiguration)
}

type ConnectorUpdateConfiguration func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateConfigurationRequest)) (*Response, error)

type ConnectorUpdateConfigurationRequest struct {
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

func (r ConnectorUpdateConfigurationRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateConfiguration) WithContext(v context.Context) func(*ConnectorUpdateConfigurationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateConfiguration) WithPretty() func(*ConnectorUpdateConfigurationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateConfiguration) WithHuman() func(*ConnectorUpdateConfigurationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateConfiguration) WithErrorTrace() func(*ConnectorUpdateConfigurationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateConfiguration) WithFilterPath(v ...string) func(*ConnectorUpdateConfigurationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateConfiguration) WithHeader(h map[string]string) func(*ConnectorUpdateConfigurationRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateConfiguration) WithOpaqueID(s string) func(*ConnectorUpdateConfigurationRequest) {
	_ = "STUB: not implemented"
	return nil
}
