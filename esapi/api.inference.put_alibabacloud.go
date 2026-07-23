package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutAlibabacloudFunc(t Transport) InferencePutAlibabacloud {
	_ = "STUB: not implemented"
	return *new(InferencePutAlibabacloud)
}

type InferencePutAlibabacloud func(body io.Reader, alibabacloud_inference_id string, task_type string, o ...func(*InferencePutAlibabacloudRequest)) (*Response, error)

type InferencePutAlibabacloudRequest struct {
	Body io.Reader

	AlibabacloudInferenceID string
	TaskType                string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutAlibabacloudRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutAlibabacloud) WithContext(v context.Context) func(*InferencePutAlibabacloudRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAlibabacloud) WithTimeout(v time.Duration) func(*InferencePutAlibabacloudRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAlibabacloud) WithPretty() func(*InferencePutAlibabacloudRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAlibabacloud) WithHuman() func(*InferencePutAlibabacloudRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAlibabacloud) WithErrorTrace() func(*InferencePutAlibabacloudRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAlibabacloud) WithFilterPath(v ...string) func(*InferencePutAlibabacloudRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAlibabacloud) WithHeader(h map[string]string) func(*InferencePutAlibabacloudRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAlibabacloud) WithOpaqueID(s string) func(*InferencePutAlibabacloudRequest) {
	_ = "STUB: not implemented"
	return nil
}
