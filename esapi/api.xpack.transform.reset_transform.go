package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTransformResetTransformFunc(t Transport) TransformResetTransform {
	_ = "STUB: not implemented"
	return *new(TransformResetTransform)
}

type TransformResetTransform func(transform_id string, o ...func(*TransformResetTransformRequest)) (*Response, error)

type TransformResetTransformRequest struct {
	TransformID string

	Force   *bool
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformResetTransformRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformResetTransform) WithContext(v context.Context) func(*TransformResetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformResetTransform) WithForce(v bool) func(*TransformResetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformResetTransform) WithTimeout(v time.Duration) func(*TransformResetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformResetTransform) WithPretty() func(*TransformResetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformResetTransform) WithHuman() func(*TransformResetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformResetTransform) WithErrorTrace() func(*TransformResetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformResetTransform) WithFilterPath(v ...string) func(*TransformResetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformResetTransform) WithHeader(h map[string]string) func(*TransformResetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformResetTransform) WithOpaqueID(s string) func(*TransformResetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}
