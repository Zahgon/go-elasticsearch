package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorLastSyncFunc(t Transport) ConnectorLastSync {
	_ = "STUB: not implemented"
	return *new(ConnectorLastSync)
}

type ConnectorLastSync func(body io.Reader, connector_id string, o ...func(*ConnectorLastSyncRequest)) (*Response, error)

type ConnectorLastSyncRequest struct {
	Body io.Reader

	ConnectorID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorLastSyncRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorLastSync) WithContext(v context.Context) func(*ConnectorLastSyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorLastSync) WithPretty() func(*ConnectorLastSyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorLastSync) WithHuman() func(*ConnectorLastSyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorLastSync) WithErrorTrace() func(*ConnectorLastSyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorLastSync) WithFilterPath(v ...string) func(*ConnectorLastSyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorLastSync) WithHeader(h map[string]string) func(*ConnectorLastSyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorLastSync) WithOpaqueID(s string) func(*ConnectorLastSyncRequest) {
	_ = "STUB: not implemented"
	return nil
}
