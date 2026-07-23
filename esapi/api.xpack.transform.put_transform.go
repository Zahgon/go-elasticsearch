package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newTransformPutTransformFunc(t Transport) TransformPutTransform {
	_ = "STUB: not implemented"
	return *new(TransformPutTransform)
}

type TransformPutTransform func(body io.Reader, transform_id string, o ...func(*TransformPutTransformRequest)) (*Response, error)

type TransformPutTransformRequest struct {
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

func (r TransformPutTransformRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformPutTransform) WithContext(v context.Context) func(*TransformPutTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPutTransform) WithDeferValidation(v bool) func(*TransformPutTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPutTransform) WithTimeout(v time.Duration) func(*TransformPutTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPutTransform) WithPretty() func(*TransformPutTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPutTransform) WithHuman() func(*TransformPutTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPutTransform) WithErrorTrace() func(*TransformPutTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPutTransform) WithFilterPath(v ...string) func(*TransformPutTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPutTransform) WithHeader(h map[string]string) func(*TransformPutTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPutTransform) WithOpaqueID(s string) func(*TransformPutTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}
