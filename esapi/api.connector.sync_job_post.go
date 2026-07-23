package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorSyncJobPostFunc(t Transport) ConnectorSyncJobPost {
	_ = "STUB: not implemented"
	return *new(ConnectorSyncJobPost)
}

type ConnectorSyncJobPost func(body io.Reader, o ...func(*ConnectorSyncJobPostRequest)) (*Response, error)

type ConnectorSyncJobPostRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSyncJobPostRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSyncJobPost) WithContext(v context.Context) func(*ConnectorSyncJobPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobPost) WithPretty() func(*ConnectorSyncJobPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobPost) WithHuman() func(*ConnectorSyncJobPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobPost) WithErrorTrace() func(*ConnectorSyncJobPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobPost) WithFilterPath(v ...string) func(*ConnectorSyncJobPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobPost) WithHeader(h map[string]string) func(*ConnectorSyncJobPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSyncJobPost) WithOpaqueID(s string) func(*ConnectorSyncJobPostRequest) {
	_ = "STUB: not implemented"
	return nil
}
