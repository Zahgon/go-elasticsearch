package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIngestDeletePipelineFunc(t Transport) IngestDeletePipeline {
	_ = "STUB: not implemented"
	return *new(IngestDeletePipeline)
}

type IngestDeletePipeline func(id string, o ...func(*IngestDeletePipelineRequest)) (*Response, error)

type IngestDeletePipelineRequest struct {
	PipelineID string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IngestDeletePipelineRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestDeletePipeline) WithContext(v context.Context) func(*IngestDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeletePipeline) WithMasterTimeout(v time.Duration) func(*IngestDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeletePipeline) WithTimeout(v time.Duration) func(*IngestDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeletePipeline) WithPretty() func(*IngestDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeletePipeline) WithHuman() func(*IngestDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeletePipeline) WithErrorTrace() func(*IngestDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeletePipeline) WithFilterPath(v ...string) func(*IngestDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeletePipeline) WithHeader(h map[string]string) func(*IngestDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeletePipeline) WithOpaqueID(s string) func(*IngestDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}
