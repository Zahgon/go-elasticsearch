package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateErrorFunc(t Transport) ConnectorUpdateError {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateError)
}

type ConnectorUpdateError func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateErrorRequest)) (*Response, error)

type ConnectorUpdateErrorRequest struct {
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

func (r ConnectorUpdateErrorRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateError) WithContext(v context.Context) func(*ConnectorUpdateErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateError) WithPretty() func(*ConnectorUpdateErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateError) WithHuman() func(*ConnectorUpdateErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateError) WithErrorTrace() func(*ConnectorUpdateErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateError) WithFilterPath(v ...string) func(*ConnectorUpdateErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateError) WithHeader(h map[string]string) func(*ConnectorUpdateErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateError) WithOpaqueID(s string) func(*ConnectorUpdateErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}
