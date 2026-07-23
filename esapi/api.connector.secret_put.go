package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorSecretPutFunc(t Transport) ConnectorSecretPut {
	_ = "STUB: not implemented"
	return *new(ConnectorSecretPut)
}

type ConnectorSecretPut func(id string, body io.Reader, o ...func(*ConnectorSecretPutRequest)) (*Response, error)

type ConnectorSecretPutRequest struct {
	DocumentID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSecretPutRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSecretPut) WithContext(v context.Context) func(*ConnectorSecretPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPut) WithPretty() func(*ConnectorSecretPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPut) WithHuman() func(*ConnectorSecretPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPut) WithErrorTrace() func(*ConnectorSecretPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPut) WithFilterPath(v ...string) func(*ConnectorSecretPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPut) WithHeader(h map[string]string) func(*ConnectorSecretPutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPut) WithOpaqueID(s string) func(*ConnectorSecretPutRequest) {
	_ = "STUB: not implemented"
	return nil
}
