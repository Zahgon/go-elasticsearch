package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTransformStopTransformFunc(t Transport) TransformStopTransform {
	_ = "STUB: not implemented"
	return *new(TransformStopTransform)
}

type TransformStopTransform func(transform_id string, o ...func(*TransformStopTransformRequest)) (*Response, error)

type TransformStopTransformRequest struct {
	TransformID string

	AllowNoMatch      *bool
	Force             *bool
	Timeout           time.Duration
	WaitForCheckpoint *bool
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformStopTransformRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformStopTransform) WithContext(v context.Context) func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithAllowNoMatch(v bool) func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithForce(v bool) func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithTimeout(v time.Duration) func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithWaitForCheckpoint(v bool) func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithWaitForCompletion(v bool) func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithPretty() func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithHuman() func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithErrorTrace() func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithFilterPath(v ...string) func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithHeader(h map[string]string) func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStopTransform) WithOpaqueID(s string) func(*TransformStopTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}
