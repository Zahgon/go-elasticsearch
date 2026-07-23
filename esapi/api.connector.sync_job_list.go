package esapi

import (
	"context"
	"net/http"
)

func newConnectorSyncJobListFunc(t Transport) ConnectorSyncJobList {
	_ = "STUB: not implemented"
	return *new(ConnectorSyncJobList)
}

type ConnectorSyncJobList func(o ...func(*ConnectorSyncJobListRequest)) (*Response, error)

type ConnectorSyncJobListRequest struct {
	ConnectorID string
	From        *int
	JobType     []string
	Size        *int
	Status      string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSyncJobListRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSyncJobList) WithContext(v context.Context) func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithConnectorID(v string) func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithFrom(v int) func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithJobType(v ...string) func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithSize(v int) func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithStatus(v string) func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithPretty() func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithHuman() func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithErrorTrace() func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithFilterPath(v ...string) func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithHeader(h map[string]string) func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobList) WithOpaqueID(s string) func(*ConnectorSyncJobListRequest) {
	_ = "STUB: not implemented"
	return nil
}
