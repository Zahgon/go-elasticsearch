package esapi

import (
	"context"
	"net/http"
)

func newLogstashGetPipelineFunc(t Transport) LogstashGetPipeline {
	_ = "STUB: not implemented"
	return *new(LogstashGetPipeline)
}

type LogstashGetPipeline func(o ...func(*LogstashGetPipelineRequest)) (*Response, error)

type LogstashGetPipelineRequest struct {
	DocumentID []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r LogstashGetPipelineRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LogstashGetPipeline) WithContext(v context.Context) func(*LogstashGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashGetPipeline) WithDocumentID(v ...string) func(*LogstashGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashGetPipeline) WithPretty() func(*LogstashGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashGetPipeline) WithHuman() func(*LogstashGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashGetPipeline) WithErrorTrace() func(*LogstashGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashGetPipeline) WithFilterPath(v ...string) func(*LogstashGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashGetPipeline) WithHeader(h map[string]string) func(*LogstashGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashGetPipeline) WithOpaqueID(s string) func(*LogstashGetPipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}
