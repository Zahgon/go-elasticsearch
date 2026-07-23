package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTransformScheduleNowTransformFunc(t Transport) TransformScheduleNowTransform {
	_ = "STUB: not implemented"
	return *new(TransformScheduleNowTransform)
}

type TransformScheduleNowTransform func(transform_id string, o ...func(*TransformScheduleNowTransformRequest)) (*Response, error)

type TransformScheduleNowTransformRequest struct {
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

func (r TransformScheduleNowTransformRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformScheduleNowTransform) WithContext(v context.Context) func(*TransformScheduleNowTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformScheduleNowTransform) WithTimeout(v time.Duration) func(*TransformScheduleNowTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformScheduleNowTransform) WithPretty() func(*TransformScheduleNowTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformScheduleNowTransform) WithHuman() func(*TransformScheduleNowTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformScheduleNowTransform) WithErrorTrace() func(*TransformScheduleNowTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformScheduleNowTransform) WithFilterPath(v ...string) func(*TransformScheduleNowTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformScheduleNowTransform) WithHeader(h map[string]string) func(*TransformScheduleNowTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformScheduleNowTransform) WithOpaqueID(s string) func(*TransformScheduleNowTransformRequest) {
	_ = "STUB: not implemented"
	return nil
}
