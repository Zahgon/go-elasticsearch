package esapi

import (
	"context"
	"net/http"
)

func newEqlGetStatusFunc(t Transport) EqlGetStatus {
	_ = "STUB: not implemented"
	return *new(EqlGetStatus)
}

type EqlGetStatus func(id string, o ...func(*EqlGetStatusRequest)) (*Response, error)

type EqlGetStatusRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EqlGetStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EqlGetStatus) WithContext(v context.Context) func(*EqlGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGetStatus) WithPretty() func(*EqlGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGetStatus) WithHuman() func(*EqlGetStatusRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlGetStatus) WithErrorTrace() func(*EqlGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGetStatus) WithFilterPath(v ...string) func(*EqlGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGetStatus) WithHeader(h map[string]string) func(*EqlGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGetStatus) WithOpaqueID(s string) func(*EqlGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
