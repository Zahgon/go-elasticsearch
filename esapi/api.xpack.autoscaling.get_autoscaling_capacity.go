package esapi

import (
	"context"
	"net/http"
	"time"
)

func newAutoscalingGetAutoscalingCapacityFunc(t Transport) AutoscalingGetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return *new(AutoscalingGetAutoscalingCapacity)
}

type AutoscalingGetAutoscalingCapacity func(o ...func(*AutoscalingGetAutoscalingCapacityRequest)) (*Response, error)

type AutoscalingGetAutoscalingCapacityRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r AutoscalingGetAutoscalingCapacityRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f AutoscalingGetAutoscalingCapacity) WithContext(v context.Context) func(*AutoscalingGetAutoscalingCapacityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingCapacity) WithMasterTimeout(v time.Duration) func(*AutoscalingGetAutoscalingCapacityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingCapacity) WithPretty() func(*AutoscalingGetAutoscalingCapacityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingCapacity) WithHuman() func(*AutoscalingGetAutoscalingCapacityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingCapacity) WithErrorTrace() func(*AutoscalingGetAutoscalingCapacityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingCapacity) WithFilterPath(v ...string) func(*AutoscalingGetAutoscalingCapacityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingCapacity) WithHeader(h map[string]string) func(*AutoscalingGetAutoscalingCapacityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AutoscalingGetAutoscalingCapacity) WithOpaqueID(s string) func(*AutoscalingGetAutoscalingCapacityRequest) {
	_ = "STUB: not implemented"
	return nil
}
