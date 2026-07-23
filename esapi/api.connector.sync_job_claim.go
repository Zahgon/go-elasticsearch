package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorSyncJobClaimFunc(t Transport) ConnectorSyncJobClaim {
	_ = "STUB: not implemented"
	return *new(ConnectorSyncJobClaim)
}

type ConnectorSyncJobClaim func(body io.Reader, connector_sync_job_id string, o ...func(*ConnectorSyncJobClaimRequest)) (*Response, error)

type ConnectorSyncJobClaimRequest struct {
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

func (r ConnectorSyncJobClaimRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSyncJobClaim) WithContext(v context.Context) func(*ConnectorSyncJobClaimRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobClaim) WithPretty() func(*ConnectorSyncJobClaimRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobClaim) WithHuman() func(*ConnectorSyncJobClaimRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobClaim) WithErrorTrace() func(*ConnectorSyncJobClaimRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobClaim) WithFilterPath(v ...string) func(*ConnectorSyncJobClaimRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobClaim) WithHeader(h map[string]string) func(*ConnectorSyncJobClaimRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobClaim) WithOpaqueID(s string) func(*ConnectorSyncJobClaimRequest) {
	_ = "STUB: not implemented"
	return nil
}
