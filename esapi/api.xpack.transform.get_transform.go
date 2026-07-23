package esapi

import (
	"context"
	"net/http"
)

func newTransformGetTransformFunc(t Transport) TransformGetTransform {
	_ = "STUB: not implemented"
	return *new(TransformGetTransform)
}

type TransformGetTransform func(o ...func(*TransformGetTransformRequest)) (*Response, error)

type TransformGetTransformRequest struct {
	TransformID []string

	AllowNoMatch     *bool
	ExcludeGenerated *bool
	From             *int
	Size             *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformGetTransformRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformGetTransform) WithContext(v context.Context) func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithTransformID(v ...string) func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithAllowNoMatch(v bool) func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithExcludeGenerated(v bool) func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithFrom(v int) func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithSize(v int) func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithPretty() func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithHuman() func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithErrorTrace() func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithFilterPath(v ...string) func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithHeader(h map[string]string) func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransform) WithOpaqueID(s string) func(*TransformGetTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}
