package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateNativeFunc(t Transport) ConnectorUpdateNative {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateNative)
}

type ConnectorUpdateNative func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateNativeRequest)) (*Response, error)

type ConnectorUpdateNativeRequest struct {
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

func (r ConnectorUpdateNativeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateNative) WithContext(v context.Context) func(*ConnectorUpdateNativeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateNative) WithPretty() func(*ConnectorUpdateNativeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateNative) WithHuman() func(*ConnectorUpdateNativeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateNative) WithErrorTrace() func(*ConnectorUpdateNativeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateNative) WithFilterPath(v ...string) func(*ConnectorUpdateNativeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateNative) WithHeader(h map[string]string) func(*ConnectorUpdateNativeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateNative) WithOpaqueID(s string) func(*ConnectorUpdateNativeRequest) {
	_ = "STUB: not implemented"
	return nil
}
