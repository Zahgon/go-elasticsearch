package esapi

import (
	"context"
	"net/http"
)

func newLogstashDeletePipelineFunc(t Transport) LogstashDeletePipeline {
	_ = "STUB: not implemented"
	return *new(LogstashDeletePipeline)
}

type LogstashDeletePipeline func(id string, o ...func(*LogstashDeletePipelineRequest)) (*Response, error)

type LogstashDeletePipelineRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r LogstashDeletePipelineRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f LogstashDeletePipeline) WithContext(v context.Context) func(*LogstashDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashDeletePipeline) WithPretty() func(*LogstashDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashDeletePipeline) WithHuman() func(*LogstashDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashDeletePipeline) WithErrorTrace() func(*LogstashDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashDeletePipeline) WithFilterPath(v ...string) func(*LogstashDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashDeletePipeline) WithHeader(h map[string]string) func(*LogstashDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f LogstashDeletePipeline) WithOpaqueID(s string) func(*LogstashDeletePipelineRequest) {
	_ = "STUB: not implemented"
	return nil
}
