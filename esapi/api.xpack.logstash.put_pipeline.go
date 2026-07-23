package esapi

import (
	"context"
	"io"
	"net/http"
)

func newLogstashPutPipelineFunc(t Transport) LogstashPutPipeline {
	_ = "STUB: not implemented"
	return *new(LogstashPutPipeline)
}

type LogstashPutPipeline func(id string, body io.Reader, o ...func(*LogstashPutPipelineRequest)) (*Response, error)

type LogstashPutPipelineRequest struct {
	DocumentID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r LogstashPutPipelineRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LogstashPutPipeline) WithContext(v context.Context) func(*LogstashPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashPutPipeline) WithPretty() func(*LogstashPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashPutPipeline) WithHuman() func(*LogstashPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashPutPipeline) WithErrorTrace() func(*LogstashPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashPutPipeline) WithFilterPath(v ...string) func(*LogstashPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashPutPipeline) WithHeader(h map[string]string) func(*LogstashPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashPutPipeline) WithOpaqueID(s string) func(*LogstashPutPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}
