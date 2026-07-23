package esapi

import (
	"context"
	"net/http"
)

func newConnectorSecretGetFunc(t Transport) ConnectorSecretGet {
	_ = "STUB: not implemented"
	return *new(ConnectorSecretGet)
}

type ConnectorSecretGet func(id string, o ...func(*ConnectorSecretGetRequest)) (*Response, error)

type ConnectorSecretGetRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSecretGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSecretGet) WithContext(v context.Context) func(*ConnectorSecretGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretGet) WithPretty() func(*ConnectorSecretGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretGet) WithHuman() func(*ConnectorSecretGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretGet) WithErrorTrace() func(*ConnectorSecretGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretGet) WithFilterPath(v ...string) func(*ConnectorSecretGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretGet) WithHeader(h map[string]string) func(*ConnectorSecretGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretGet) WithOpaqueID(s string) func(*ConnectorSecretGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
