package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorPutFunc(t Transport) ConnectorPut {
	_ = "STUB: not implemented"
	return *new(ConnectorPut)
}

type ConnectorPut func(o ...func(*ConnectorPutRequest)) (*Response, error)

type ConnectorPutRequest struct {
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

func (r ConnectorPutRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorPut) WithContext(v context.Context) func(*ConnectorPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPut) WithBody(v io.Reader) func(*ConnectorPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPut) WithConnectorID(v string) func(*ConnectorPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPut) WithPretty() func(*ConnectorPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPut) WithHuman() func(*ConnectorPutRequest) { _ = "STUB: not implemented"; return nil }

func (f ConnectorPut) WithErrorTrace() func(*ConnectorPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPut) WithFilterPath(v ...string) func(*ConnectorPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPut) WithHeader(h map[string]string) func(*ConnectorPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorPut) WithOpaqueID(s string) func(*ConnectorPutRequest) {
	_ = "STUB: not implemented"
	return nil
}
