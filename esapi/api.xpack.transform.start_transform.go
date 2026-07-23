package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTransformStartTransformFunc(t Transport) TransformStartTransform {
	_ = "STUB: not implemented"
	return *new(TransformStartTransform)
}

type TransformStartTransform func(transform_id string, o ...func(*TransformStartTransformRequest)) (*Response, error)

type TransformStartTransformRequest struct {
	TransformID string

	From    string
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformStartTransformRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformStartTransform) WithContext(v context.Context) func(*TransformStartTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStartTransform) WithFrom(v string) func(*TransformStartTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStartTransform) WithTimeout(v time.Duration) func(*TransformStartTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStartTransform) WithPretty() func(*TransformStartTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStartTransform) WithHuman() func(*TransformStartTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStartTransform) WithErrorTrace() func(*TransformStartTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStartTransform) WithFilterPath(v ...string) func(*TransformStartTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStartTransform) WithHeader(h map[string]string) func(*TransformStartTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformStartTransform) WithOpaqueID(s string) func(*TransformStartTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}
