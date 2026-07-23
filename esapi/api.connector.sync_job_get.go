package esapi

import (
	"context"
	"net/http"
)

func newConnectorSyncJobGetFunc(t Transport) ConnectorSyncJobGet {
	_ = "STUB: not implemented"
	return *new(ConnectorSyncJobGet)
}

type ConnectorSyncJobGet func(connector_sync_job_id string, o ...func(*ConnectorSyncJobGetRequest)) (*Response, error)

type ConnectorSyncJobGetRequest struct {
	ConnectorSyncJobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSyncJobGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSyncJobGet) WithContext(v context.Context) func(*ConnectorSyncJobGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobGet) WithPretty() func(*ConnectorSyncJobGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobGet) WithHuman() func(*ConnectorSyncJobGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobGet) WithErrorTrace() func(*ConnectorSyncJobGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobGet) WithFilterPath(v ...string) func(*ConnectorSyncJobGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobGet) WithHeader(h map[string]string) func(*ConnectorSyncJobGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobGet) WithOpaqueID(s string) func(*ConnectorSyncJobGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
