package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateSchedulingFunc(t Transport) ConnectorUpdateScheduling {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateScheduling)
}

type ConnectorUpdateScheduling func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateSchedulingRequest)) (*Response, error)

type ConnectorUpdateSchedulingRequest struct {
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

func (r ConnectorUpdateSchedulingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateScheduling) WithContext(v context.Context) func(*ConnectorUpdateSchedulingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateScheduling) WithPretty() func(*ConnectorUpdateSchedulingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateScheduling) WithHuman() func(*ConnectorUpdateSchedulingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateScheduling) WithErrorTrace() func(*ConnectorUpdateSchedulingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateScheduling) WithFilterPath(v ...string) func(*ConnectorUpdateSchedulingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateScheduling) WithHeader(h map[string]string) func(*ConnectorUpdateSchedulingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateScheduling) WithOpaqueID(s string) func(*ConnectorUpdateSchedulingRequest) {
	_ = "STUB: not implemented"
	return nil
}
