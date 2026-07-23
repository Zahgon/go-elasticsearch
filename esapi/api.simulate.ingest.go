package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSimulateIngestFunc(t Transport) SimulateIngest {
	_ = "STUB: not implemented"
	return *new(SimulateIngest)
}

type SimulateIngest func(body io.Reader, o ...func(*SimulateIngestRequest)) (*Response, error)

type SimulateIngestRequest struct {
	Index string

	Body io.Reader

	MergeType string
	Pipeline  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SimulateIngestRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SimulateIngest) WithContext(v context.Context) func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SimulateIngest) WithIndex(v string) func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SimulateIngest) WithMergeType(v string) func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SimulateIngest) WithPipeline(v string) func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SimulateIngest) WithPretty() func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SimulateIngest) WithHuman() func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SimulateIngest) WithErrorTrace() func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SimulateIngest) WithFilterPath(v ...string) func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SimulateIngest) WithHeader(h map[string]string) func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SimulateIngest) WithOpaqueID(s string) func(*SimulateIngestRequest) {
	_ = "STUB: not implemented"
	return nil
}
