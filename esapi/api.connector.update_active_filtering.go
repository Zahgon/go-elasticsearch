package esapi

import (
	"context"
	"net/http"
)

func newConnectorUpdateActiveFilteringFunc(t Transport) ConnectorUpdateActiveFiltering {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateActiveFiltering)
}

type ConnectorUpdateActiveFiltering func(connector_id string, o ...func(*ConnectorUpdateActiveFilteringRequest)) (*Response, error)

type ConnectorUpdateActiveFilteringRequest struct {
	ConnectorID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorUpdateActiveFilteringRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateActiveFiltering) WithContext(v context.Context) func(*ConnectorUpdateActiveFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateActiveFiltering) WithPretty() func(*ConnectorUpdateActiveFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateActiveFiltering) WithHuman() func(*ConnectorUpdateActiveFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateActiveFiltering) WithErrorTrace() func(*ConnectorUpdateActiveFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateActiveFiltering) WithFilterPath(v ...string) func(*ConnectorUpdateActiveFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateActiveFiltering) WithHeader(h map[string]string) func(*ConnectorUpdateActiveFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateActiveFiltering) WithOpaqueID(s string) func(*ConnectorUpdateActiveFilteringRequest) {
	_ = "STUB: not implemented"
	return nil
}
