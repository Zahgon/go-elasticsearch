package esapi

import (
	"context"
	"net/http"
)

func newConnectorSyncJobCancelFunc(t Transport) ConnectorSyncJobCancel {
	_ = "STUB: not implemented"
	return *new(ConnectorSyncJobCancel)
}

type ConnectorSyncJobCancel func(connector_sync_job_id string, o ...func(*ConnectorSyncJobCancelRequest)) (*Response, error)

type ConnectorSyncJobCancelRequest struct {
	ConnectorSyncJobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSyncJobCancelRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSyncJobCancel) WithContext(v context.Context) func(*ConnectorSyncJobCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCancel) WithPretty() func(*ConnectorSyncJobCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCancel) WithHuman() func(*ConnectorSyncJobCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCancel) WithErrorTrace() func(*ConnectorSyncJobCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCancel) WithFilterPath(v ...string) func(*ConnectorSyncJobCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCancel) WithHeader(h map[string]string) func(*ConnectorSyncJobCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCancel) WithOpaqueID(s string) func(*ConnectorSyncJobCancelRequest) {
	_ = "STUB: not implemented"
	return nil
}
