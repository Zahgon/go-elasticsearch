package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIngestGetPipelineFunc(t Transport) IngestGetPipeline {
	_ = "STUB: not implemented"
	return *new(IngestGetPipeline)
}

type IngestGetPipeline func(o ...func(*IngestGetPipelineRequest)) (*Response, error)

type IngestGetPipelineRequest struct {
	PipelineID string

	MasterTimeout time.Duration
	Summary       *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IngestGetPipelineRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestGetPipeline) WithContext(v context.Context) func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetPipeline) WithPipelineID(v string) func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetPipeline) WithMasterTimeout(v time.Duration) func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetPipeline) WithSummary(v bool) func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetPipeline) WithPretty() func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetPipeline) WithHuman() func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetPipeline) WithErrorTrace() func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetPipeline) WithFilterPath(v ...string) func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetPipeline) WithHeader(h map[string]string) func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetPipeline) WithOpaqueID(s string) func(*IngestGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}
