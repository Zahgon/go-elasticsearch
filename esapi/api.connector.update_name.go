package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateNameFunc(t Transport) ConnectorUpdateName {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateName)
}

type ConnectorUpdateName func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateNameRequest)) (*Response, error)

type ConnectorUpdateNameRequest struct {
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

func (r ConnectorUpdateNameRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateName) WithContext(v context.Context) func(*ConnectorUpdateNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateName) WithPretty() func(*ConnectorUpdateNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateName) WithHuman() func(*ConnectorUpdateNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateName) WithErrorTrace() func(*ConnectorUpdateNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateName) WithFilterPath(v ...string) func(*ConnectorUpdateNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateName) WithHeader(h map[string]string) func(*ConnectorUpdateNameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateName) WithOpaqueID(s string) func(*ConnectorUpdateNameRequest) {
	_ = "STUB: not implemented"
	return nil
}
