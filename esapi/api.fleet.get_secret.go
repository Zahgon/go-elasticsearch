package esapi

import (
	"context"
	"net/http"
)

func newFleetGetSecretFunc(t Transport) FleetGetSecret {
	_ = "STUB: not implemented"
	return *new(FleetGetSecret)
}

type FleetGetSecret func(id string, o ...func(*FleetGetSecretRequest)) (*Response, error)

type FleetGetSecretRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r FleetGetSecretRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FleetGetSecret) WithContext(v context.Context) func(*FleetGetSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGetSecret) WithPretty() func(*FleetGetSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGetSecret) WithHuman() func(*FleetGetSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGetSecret) WithErrorTrace() func(*FleetGetSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGetSecret) WithFilterPath(v ...string) func(*FleetGetSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGetSecret) WithHeader(h map[string]string) func(*FleetGetSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetGetSecret) WithOpaqueID(s string) func(*FleetGetSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}
