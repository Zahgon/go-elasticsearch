package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newAutoscalingPutAutoscalingPolicyFunc(t Transport) AutoscalingPutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return *new(AutoscalingPutAutoscalingPolicy)
}

type AutoscalingPutAutoscalingPolicy func(name string, body io.Reader, o ...func(*AutoscalingPutAutoscalingPolicyRequest)) (*Response, error)

type AutoscalingPutAutoscalingPolicyRequest struct {
	Body io.Reader

	Name string

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

func (r AutoscalingPutAutoscalingPolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f AutoscalingPutAutoscalingPolicy) WithContext(v context.Context) func(*AutoscalingPutAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingPutAutoscalingPolicy) WithMasterTimeout(v time.Duration) func(*AutoscalingPutAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingPutAutoscalingPolicy) WithTimeout(v time.Duration) func(*AutoscalingPutAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingPutAutoscalingPolicy) WithPretty() func(*AutoscalingPutAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingPutAutoscalingPolicy) WithHuman() func(*AutoscalingPutAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingPutAutoscalingPolicy) WithErrorTrace() func(*AutoscalingPutAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingPutAutoscalingPolicy) WithFilterPath(v ...string) func(*AutoscalingPutAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingPutAutoscalingPolicy) WithHeader(h map[string]string) func(*AutoscalingPutAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingPutAutoscalingPolicy) WithOpaqueID(s string) func(*AutoscalingPutAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}
