package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newTransformPreviewTransformFunc(t Transport) TransformPreviewTransform {
	_ = "STUB: not implemented"
	return *new(TransformPreviewTransform)
}

type TransformPreviewTransform func(o ...func(*TransformPreviewTransformRequest)) (*Response, error)

type TransformPreviewTransformRequest struct {
	Body io.Reader

	TransformID string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformPreviewTransformRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformPreviewTransform) WithContext(v context.Context) func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPreviewTransform) WithBody(v io.Reader) func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPreviewTransform) WithTransformID(v string) func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPreviewTransform) WithTimeout(v time.Duration) func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPreviewTransform) WithPretty() func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPreviewTransform) WithHuman() func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPreviewTransform) WithErrorTrace() func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPreviewTransform) WithFilterPath(v ...string) func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPreviewTransform) WithHeader(h map[string]string) func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformPreviewTransform) WithOpaqueID(s string) func(*TransformPreviewTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}
