package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newTransformUpdateTransformFunc(t Transport) TransformUpdateTransform {
	_ = "STUB: not implemented"
	return *new(TransformUpdateTransform)
}

type TransformUpdateTransform func(body io.Reader, transform_id string, o ...func(*TransformUpdateTransformRequest)) (*Response, error)

type TransformUpdateTransformRequest struct {
	Body io.Reader

	TransformID string

	DeferValidation *bool
	Timeout         time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformUpdateTransformRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformUpdateTransform) WithContext(v context.Context) func(*TransformUpdateTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpdateTransform) WithDeferValidation(v bool) func(*TransformUpdateTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpdateTransform) WithTimeout(v time.Duration) func(*TransformUpdateTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpdateTransform) WithPretty() func(*TransformUpdateTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpdateTransform) WithHuman() func(*TransformUpdateTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpdateTransform) WithErrorTrace() func(*TransformUpdateTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpdateTransform) WithFilterPath(v ...string) func(*TransformUpdateTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpdateTransform) WithHeader(h map[string]string) func(*TransformUpdateTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpdateTransform) WithOpaqueID(s string) func(*TransformUpdateTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}
