package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorSyncJobErrorFunc(t Transport) ConnectorSyncJobError {
	_ = "STUB: not implemented"
	return *new(ConnectorSyncJobError)
}

type ConnectorSyncJobError func(body io.Reader, connector_sync_job_id string, o ...func(*ConnectorSyncJobErrorRequest)) (*Response, error)

type ConnectorSyncJobErrorRequest struct {
	Body io.Reader

	ConnectorSyncJobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSyncJobErrorRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSyncJobError) WithContext(v context.Context) func(*ConnectorSyncJobErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobError) WithPretty() func(*ConnectorSyncJobErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobError) WithHuman() func(*ConnectorSyncJobErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobError) WithErrorTrace() func(*ConnectorSyncJobErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobError) WithFilterPath(v ...string) func(*ConnectorSyncJobErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobError) WithHeader(h map[string]string) func(*ConnectorSyncJobErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobError) WithOpaqueID(s string) func(*ConnectorSyncJobErrorRequest) {
	_ = "STUB: not implemented"
	return nil
}
