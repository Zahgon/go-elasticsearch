package esapi

import (
	"context"
	"io"
	"net/http"
)

func newIngestSimulateFunc(t Transport) IngestSimulate {
	_ = "STUB: not implemented"
	return *new(IngestSimulate)
}

type IngestSimulate func(body io.Reader, o ...func(*IngestSimulateRequest)) (*Response, error)

type IngestSimulateRequest struct {
	PipelineID string

	Body io.Reader

	Verbose *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IngestSimulateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestSimulate) WithContext(v context.Context) func(*IngestSimulateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestSimulate) WithPipelineID(v string) func(*IngestSimulateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestSimulate) WithVerbose(v bool) func(*IngestSimulateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestSimulate) WithPretty() func(*IngestSimulateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestSimulate) WithHuman() func(*IngestSimulateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestSimulate) WithErrorTrace() func(*IngestSimulateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestSimulate) WithFilterPath(v ...string) func(*IngestSimulateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestSimulate) WithHeader(h map[string]string) func(*IngestSimulateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestSimulate) WithOpaqueID(s string) func(*IngestSimulateRequest) {
	_ = "STUB: not implemented"
	return nil
}
