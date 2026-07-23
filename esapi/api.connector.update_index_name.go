package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateIndexNameFunc(t Transport) ConnectorUpdateIndexName {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateIndexName)
}

type ConnectorUpdateIndexName func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateIndexNameRequest)) (*Response, error)

type ConnectorUpdateIndexNameRequest struct {
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

func (r ConnectorUpdateIndexNameRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateIndexName) WithContext(v context.Context) func(*ConnectorUpdateIndexNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateIndexName) WithPretty() func(*ConnectorUpdateIndexNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateIndexName) WithHuman() func(*ConnectorUpdateIndexNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateIndexName) WithErrorTrace() func(*ConnectorUpdateIndexNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateIndexName) WithFilterPath(v ...string) func(*ConnectorUpdateIndexNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateIndexName) WithHeader(h map[string]string) func(*ConnectorUpdateIndexNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateIndexName) WithOpaqueID(s string) func(*ConnectorUpdateIndexNameRequest) {
	_ = "STUB: not implemented"
	return nil
}
