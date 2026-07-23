package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutElasticsearchFunc(t Transport) InferencePutElasticsearch {
	_ = "STUB: not implemented"
	return *new(InferencePutElasticsearch)
}

type InferencePutElasticsearch func(body io.Reader, elasticsearch_inference_id string, task_type string, o ...func(*InferencePutElasticsearchRequest)) (*Response, error)

type InferencePutElasticsearchRequest struct {
	Body io.Reader

	ElasticsearchInferenceID string
	TaskType                 string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutElasticsearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutElasticsearch) WithContext(v context.Context) func(*InferencePutElasticsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElasticsearch) WithTimeout(v time.Duration) func(*InferencePutElasticsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElasticsearch) WithPretty() func(*InferencePutElasticsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElasticsearch) WithHuman() func(*InferencePutElasticsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElasticsearch) WithErrorTrace() func(*InferencePutElasticsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElasticsearch) WithFilterPath(v ...string) func(*InferencePutElasticsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElasticsearch) WithHeader(h map[string]string) func(*InferencePutElasticsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElasticsearch) WithOpaqueID(s string) func(*InferencePutElasticsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}
