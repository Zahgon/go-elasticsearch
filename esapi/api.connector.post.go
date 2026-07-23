package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorPostFunc(t Transport) ConnectorPost {
	_ = "STUB: not implemented"
	return *new(ConnectorPost)
}

type ConnectorPost func(o ...func(*ConnectorPostRequest)) (*Response, error)

type ConnectorPostRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorPostRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorPost) WithContext(v context.Context) func(*ConnectorPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPost) WithBody(v io.Reader) func(*ConnectorPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPost) WithPretty() func(*ConnectorPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPost) WithHuman() func(*ConnectorPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPost) WithErrorTrace() func(*ConnectorPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPost) WithFilterPath(v ...string) func(*ConnectorPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPost) WithHeader(h map[string]string) func(*ConnectorPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPost) WithOpaqueID(s string) func(*ConnectorPostRequest) {
	_ = "STUB: not implemented"
	return nil
}
