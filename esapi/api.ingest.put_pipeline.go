package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIngestPutPipelineFunc(t Transport) IngestPutPipeline {
	_ = "STUB: not implemented"
	return *new(IngestPutPipeline)
}

type IngestPutPipeline func(id string, body io.Reader, o ...func(*IngestPutPipelineRequest)) (*Response, error)

type IngestPutPipelineRequest struct {
	PipelineID string

	Body io.Reader

	IfVersion     *int
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

func (r IngestPutPipelineRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestPutPipeline) WithContext(v context.Context) func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutPipeline) WithIfVersion(v int) func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutPipeline) WithMasterTimeout(v time.Duration) func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutPipeline) WithTimeout(v time.Duration) func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutPipeline) WithPretty() func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutPipeline) WithHuman() func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutPipeline) WithErrorTrace() func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutPipeline) WithFilterPath(v ...string) func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutPipeline) WithHeader(h map[string]string) func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutPipeline) WithOpaqueID(s string) func(*IngestPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}
