package esapi

import (
	"context"
	"net/http"
)

func newConnectorCheckInFunc(t Transport) ConnectorCheckIn {
	_ = "STUB: not implemented"
	return *new(ConnectorCheckIn)
}

type ConnectorCheckIn func(connector_id string, o ...func(*ConnectorCheckInRequest)) (*Response, error)

type ConnectorCheckInRequest struct {
	ConnectorID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorCheckInRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorCheckIn) WithContext(v context.Context) func(*ConnectorCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorCheckIn) WithPretty() func(*ConnectorCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorCheckIn) WithHuman() func(*ConnectorCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorCheckIn) WithErrorTrace() func(*ConnectorCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorCheckIn) WithFilterPath(v ...string) func(*ConnectorCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorCheckIn) WithHeader(h map[string]string) func(*ConnectorCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorCheckIn) WithOpaqueID(s string) func(*ConnectorCheckInRequest) {
	_ = "STUB: not implemented"
	return nil
}
