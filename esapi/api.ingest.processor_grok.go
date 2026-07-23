package esapi

import (
	"context"
	"net/http"
)

func newIngestProcessorGrokFunc(t Transport) IngestProcessorGrok {
	_ = "STUB: not implemented"
	return *new(IngestProcessorGrok)
}

type IngestProcessorGrok func(o ...func(*IngestProcessorGrokRequest)) (*Response, error)

type IngestProcessorGrokRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IngestProcessorGrokRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestProcessorGrok) WithContext(v context.Context) func(*IngestProcessorGrokRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestProcessorGrok) WithPretty() func(*IngestProcessorGrokRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestProcessorGrok) WithHuman() func(*IngestProcessorGrokRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestProcessorGrok) WithErrorTrace() func(*IngestProcessorGrokRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestProcessorGrok) WithFilterPath(v ...string) func(*IngestProcessorGrokRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestProcessorGrok) WithHeader(h map[string]string) func(*IngestProcessorGrokRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestProcessorGrok) WithOpaqueID(s string) func(*IngestProcessorGrokRequest) {
	_ = "STUB: not implemented"
	return nil
}
