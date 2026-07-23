package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateStatusFunc(t Transport) ConnectorUpdateStatus {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateStatus)
}

type ConnectorUpdateStatus func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateStatusRequest)) (*Response, error)

type ConnectorUpdateStatusRequest struct {
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

func (r ConnectorUpdateStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateStatus) WithContext(v context.Context) func(*ConnectorUpdateStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateStatus) WithPretty() func(*ConnectorUpdateStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateStatus) WithHuman() func(*ConnectorUpdateStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateStatus) WithErrorTrace() func(*ConnectorUpdateStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateStatus) WithFilterPath(v ...string) func(*ConnectorUpdateStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateStatus) WithHeader(h map[string]string) func(*ConnectorUpdateStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateStatus) WithOpaqueID(s string) func(*ConnectorUpdateStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
