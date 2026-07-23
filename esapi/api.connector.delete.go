package esapi

import (
	"context"
	"net/http"
)

func newConnectorDeleteFunc(t Transport) ConnectorDelete {
	_ = "STUB: not implemented"
	return *new(ConnectorDelete)
}

type ConnectorDelete func(connector_id string, o ...func(*ConnectorDeleteRequest)) (*Response, error)

type ConnectorDeleteRequest struct {
	ConnectorID string

	DeleteSyncJobs *bool
	Hard           *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorDelete) WithContext(v context.Context) func(*ConnectorDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorDelete) WithDeleteSyncJobs(v bool) func(*ConnectorDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorDelete) WithHard(v bool) func(*ConnectorDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorDelete) WithPretty() func(*ConnectorDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorDelete) WithHuman() func(*ConnectorDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorDelete) WithErrorTrace() func(*ConnectorDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorDelete) WithFilterPath(v ...string) func(*ConnectorDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorDelete) WithHeader(h map[string]string) func(*ConnectorDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorDelete) WithOpaqueID(s string) func(*ConnectorDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
