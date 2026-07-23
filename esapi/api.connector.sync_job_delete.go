package esapi

import (
	"context"
	"net/http"
)

func newConnectorSyncJobDeleteFunc(t Transport) ConnectorSyncJobDelete {
	_ = "STUB: not implemented"
	return *new(ConnectorSyncJobDelete)
}

type ConnectorSyncJobDelete func(connector_sync_job_id string, o ...func(*ConnectorSyncJobDeleteRequest)) (*Response, error)

type ConnectorSyncJobDeleteRequest struct {
	ConnectorSyncJobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSyncJobDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSyncJobDelete) WithContext(v context.Context) func(*ConnectorSyncJobDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobDelete) WithPretty() func(*ConnectorSyncJobDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobDelete) WithHuman() func(*ConnectorSyncJobDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobDelete) WithErrorTrace() func(*ConnectorSyncJobDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobDelete) WithFilterPath(v ...string) func(*ConnectorSyncJobDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobDelete) WithHeader(h map[string]string) func(*ConnectorSyncJobDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobDelete) WithOpaqueID(s string) func(*ConnectorSyncJobDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
