package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorSyncJobUpdateStatsFunc(t Transport) ConnectorSyncJobUpdateStats {
	_ = "STUB: not implemented"
	return *new(ConnectorSyncJobUpdateStats)
}

type ConnectorSyncJobUpdateStats func(body io.Reader, connector_sync_job_id string, o ...func(*ConnectorSyncJobUpdateStatsRequest)) (*Response, error)

type ConnectorSyncJobUpdateStatsRequest struct {
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

func (r ConnectorSyncJobUpdateStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSyncJobUpdateStats) WithContext(v context.Context) func(*ConnectorSyncJobUpdateStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobUpdateStats) WithPretty() func(*ConnectorSyncJobUpdateStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobUpdateStats) WithHuman() func(*ConnectorSyncJobUpdateStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobUpdateStats) WithErrorTrace() func(*ConnectorSyncJobUpdateStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobUpdateStats) WithFilterPath(v ...string) func(*ConnectorSyncJobUpdateStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobUpdateStats) WithHeader(h map[string]string) func(*ConnectorSyncJobUpdateStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobUpdateStats) WithOpaqueID(s string) func(*ConnectorSyncJobUpdateStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
