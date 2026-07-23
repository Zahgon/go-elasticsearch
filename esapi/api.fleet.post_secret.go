package esapi

import (
	"context"
	"io"
	"net/http"
)

func newFleetPostSecretFunc(t Transport) FleetPostSecret {
	_ = "STUB: not implemented"
	return *new(FleetPostSecret)
}

type FleetPostSecret func(body io.Reader, o ...func(*FleetPostSecretRequest)) (*Response, error)

type FleetPostSecretRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r FleetPostSecretRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FleetPostSecret) WithContext(v context.Context) func(*FleetPostSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetPostSecret) WithPretty() func(*FleetPostSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetPostSecret) WithHuman() func(*FleetPostSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetPostSecret) WithErrorTrace() func(*FleetPostSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetPostSecret) WithFilterPath(v ...string) func(*FleetPostSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetPostSecret) WithHeader(h map[string]string) func(*FleetPostSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetPostSecret) WithOpaqueID(s string) func(*FleetPostSecretRequest) {
	_ = "STUB: not implemented"
	return nil
}
