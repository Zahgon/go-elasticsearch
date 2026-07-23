package esapi

import (
	"context"
	"net/http"
)

func newConnectorSyncJobCheckInFunc(t Transport) ConnectorSyncJobCheckIn {
	_ = "STUB: not implemented"
	return *new(ConnectorSyncJobCheckIn)
}

type ConnectorSyncJobCheckIn func(connector_sync_job_id string, o ...func(*ConnectorSyncJobCheckInRequest)) (*Response, error)

type ConnectorSyncJobCheckInRequest struct {
	ConnectorSyncJobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSyncJobCheckInRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSyncJobCheckIn) WithContext(v context.Context) func(*ConnectorSyncJobCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCheckIn) WithPretty() func(*ConnectorSyncJobCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCheckIn) WithHuman() func(*ConnectorSyncJobCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCheckIn) WithErrorTrace() func(*ConnectorSyncJobCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCheckIn) WithFilterPath(v ...string) func(*ConnectorSyncJobCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCheckIn) WithHeader(h map[string]string) func(*ConnectorSyncJobCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobCheckIn) WithOpaqueID(s string) func(*ConnectorSyncJobCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}
