package esapi

import (
	"context"
	"net/http"
	"time"
)

func newAutoscalingDeleteAutoscalingPolicyFunc(t Transport) AutoscalingDeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return *new(AutoscalingDeleteAutoscalingPolicy)
}

type AutoscalingDeleteAutoscalingPolicy func(name string, o ...func(*AutoscalingDeleteAutoscalingPolicyRequest)) (*Response, error)

type AutoscalingDeleteAutoscalingPolicyRequest struct {
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

func (r AutoscalingDeleteAutoscalingPolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f AutoscalingDeleteAutoscalingPolicy) WithContext(v context.Context) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingDeleteAutoscalingPolicy) WithMasterTimeout(v time.Duration) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingDeleteAutoscalingPolicy) WithTimeout(v time.Duration) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingDeleteAutoscalingPolicy) WithPretty() func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingDeleteAutoscalingPolicy) WithHuman() func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingDeleteAutoscalingPolicy) WithErrorTrace() func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingDeleteAutoscalingPolicy) WithFilterPath(v ...string) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingDeleteAutoscalingPolicy) WithHeader(h map[string]string) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingDeleteAutoscalingPolicy) WithOpaqueID(s string) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}
