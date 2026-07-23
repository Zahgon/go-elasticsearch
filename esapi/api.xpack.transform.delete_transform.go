package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTransformDeleteTransformFunc(t Transport) TransformDeleteTransform {
	_ = "STUB: not implemented"
	return *new(TransformDeleteTransform)
}

type TransformDeleteTransform func(transform_id string, o ...func(*TransformDeleteTransformRequest)) (*Response, error)

type TransformDeleteTransformRequest struct {
	TransformID string

	DeleteDestIndex *bool
	Force           *bool
	Timeout         time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformDeleteTransformRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformDeleteTransform) WithContext(v context.Context) func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformDeleteTransform) WithDeleteDestIndex(v bool) func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformDeleteTransform) WithForce(v bool) func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformDeleteTransform) WithTimeout(v time.Duration) func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformDeleteTransform) WithPretty() func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformDeleteTransform) WithHuman() func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformDeleteTransform) WithErrorTrace() func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformDeleteTransform) WithFilterPath(v ...string) func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformDeleteTransform) WithHeader(h map[string]string) func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformDeleteTransform) WithOpaqueID(s string) func(*TransformDeleteTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}
