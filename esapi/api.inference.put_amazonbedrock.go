package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutAmazonbedrockFunc(t Transport) InferencePutAmazonbedrock {
	_ = "STUB: not implemented"
	return *new(InferencePutAmazonbedrock)
}

type InferencePutAmazonbedrock func(body io.Reader, amazonbedrock_inference_id string, task_type string, o ...func(*InferencePutAmazonbedrockRequest)) (*Response, error)

type InferencePutAmazonbedrockRequest struct {
	Body io.Reader

	AmazonbedrockInferenceID string
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

func (r InferencePutAmazonbedrockRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutAmazonbedrock) WithContext(v context.Context) func(*InferencePutAmazonbedrockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonbedrock) WithTimeout(v time.Duration) func(*InferencePutAmazonbedrockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonbedrock) WithPretty() func(*InferencePutAmazonbedrockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonbedrock) WithHuman() func(*InferencePutAmazonbedrockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonbedrock) WithErrorTrace() func(*InferencePutAmazonbedrockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonbedrock) WithFilterPath(v ...string) func(*InferencePutAmazonbedrockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonbedrock) WithHeader(h map[string]string) func(*InferencePutAmazonbedrockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonbedrock) WithOpaqueID(s string) func(*InferencePutAmazonbedrockRequest) {
	_ = "STUB: not implemented"
	return nil
}
