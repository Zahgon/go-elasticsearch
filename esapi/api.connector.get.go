package esapi

import (
	"context"
	"net/http"
)

func newConnectorGetFunc(t Transport) ConnectorGet {
	_ = "STUB: not implemented"
	return *new(ConnectorGet)
}

type ConnectorGet func(connector_id string, o ...func(*ConnectorGetRequest)) (*Response, error)

type ConnectorGetRequest struct {
	ConnectorID string

	IncludeDeleted *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorGet) WithContext(v context.Context) func(*ConnectorGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorGet) WithIncludeDeleted(v bool) func(*ConnectorGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorGet) WithPretty() func(*ConnectorGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorGet) WithHuman() func(*ConnectorGetRequest) { _ = "STUB: not implemented"; return nil }

func (f ConnectorGet) WithErrorTrace() func(*ConnectorGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorGet) WithFilterPath(v ...string) func(*ConnectorGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorGet) WithHeader(h map[string]string) func(*ConnectorGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorGet) WithOpaqueID(s string) func(*ConnectorGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
