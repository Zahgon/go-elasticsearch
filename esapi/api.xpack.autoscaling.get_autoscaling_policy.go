package esapi

import (
	"context"
	"net/http"
	"time"
)

func newAutoscalingGetAutoscalingPolicyFunc(t Transport) AutoscalingGetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return *new(AutoscalingGetAutoscalingPolicy)
}

type AutoscalingGetAutoscalingPolicy func(name string, o ...func(*AutoscalingGetAutoscalingPolicyRequest)) (*Response, error)

type AutoscalingGetAutoscalingPolicyRequest struct {
	Name string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r AutoscalingGetAutoscalingPolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f AutoscalingGetAutoscalingPolicy) WithContext(v context.Context) func(*AutoscalingGetAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingPolicy) WithMasterTimeout(v time.Duration) func(*AutoscalingGetAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingPolicy) WithPretty() func(*AutoscalingGetAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingPolicy) WithHuman() func(*AutoscalingGetAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingPolicy) WithErrorTrace() func(*AutoscalingGetAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingPolicy) WithFilterPath(v ...string) func(*AutoscalingGetAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingPolicy) WithHeader(h map[string]string) func(*AutoscalingGetAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingPolicy) WithOpaqueID(s string) func(*AutoscalingGetAutoscalingPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}
