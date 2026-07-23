package esapi

import (
	"context"
	"net/http"
)

func newFleetDeleteSecretFunc(t Transport) FleetDeleteSecret {
	_ = "STUB: not implemented"
	return *new(FleetDeleteSecret)
}

type FleetDeleteSecret func(id string, o ...func(*FleetDeleteSecretRequest)) (*Response, error)

type FleetDeleteSecretRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r FleetDeleteSecretRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FleetDeleteSecret) WithContext(v context.Context) func(*FleetDeleteSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetDeleteSecret) WithPretty() func(*FleetDeleteSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetDeleteSecret) WithHuman() func(*FleetDeleteSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetDeleteSecret) WithErrorTrace() func(*FleetDeleteSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetDeleteSecret) WithFilterPath(v ...string) func(*FleetDeleteSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetDeleteSecret) WithHeader(h map[string]string) func(*FleetDeleteSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetDeleteSecret) WithOpaqueID(s string) func(*FleetDeleteSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}
