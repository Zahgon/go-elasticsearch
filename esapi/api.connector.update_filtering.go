package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateFilteringFunc(t Transport) ConnectorUpdateFiltering {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateFiltering)
}

type ConnectorUpdateFiltering func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateFilteringRequest)) (*Response, error)

type ConnectorUpdateFilteringRequest struct {
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

func (r ConnectorUpdateFilteringRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateFiltering) WithContext(v context.Context) func(*ConnectorUpdateFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFiltering) WithPretty() func(*ConnectorUpdateFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFiltering) WithHuman() func(*ConnectorUpdateFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFiltering) WithErrorTrace() func(*ConnectorUpdateFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFiltering) WithFilterPath(v ...string) func(*ConnectorUpdateFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFiltering) WithHeader(h map[string]string) func(*ConnectorUpdateFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateFiltering) WithOpaqueID(s string) func(*ConnectorUpdateFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}
