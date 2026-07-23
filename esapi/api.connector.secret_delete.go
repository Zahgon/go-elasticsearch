package esapi

import (
	"context"
	"net/http"
)

func newConnectorSecretDeleteFunc(t Transport) ConnectorSecretDelete {
	_ = "STUB: not implemented"
	return *new(ConnectorSecretDelete)
}

type ConnectorSecretDelete func(id string, o ...func(*ConnectorSecretDeleteRequest)) (*Response, error)

type ConnectorSecretDeleteRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSecretDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSecretDelete) WithContext(v context.Context) func(*ConnectorSecretDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretDelete) WithPretty() func(*ConnectorSecretDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretDelete) WithHuman() func(*ConnectorSecretDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretDelete) WithErrorTrace() func(*ConnectorSecretDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretDelete) WithFilterPath(v ...string) func(*ConnectorSecretDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretDelete) WithHeader(h map[string]string) func(*ConnectorSecretDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretDelete) WithOpaqueID(s string) func(*ConnectorSecretDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
