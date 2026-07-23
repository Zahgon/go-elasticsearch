package esapi

import (
	"context"
	"net/http"
)

func newMLGetDatafeedsFunc(t Transport) MLGetDatafeeds {
	_ = "STUB: not implemented"
	return *new(MLGetDatafeeds)
}

type MLGetDatafeeds func(o ...func(*MLGetDatafeedsRequest)) (*Response, error)

type MLGetDatafeedsRequest struct {
	DatafeedID []string

	AllowNoMatch     *bool
	ExcludeGenerated *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetDatafeedsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetDatafeeds) WithContext(v context.Context) func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeeds) WithDatafeedID(v ...string) func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeeds) WithAllowNoMatch(v bool) func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeeds) WithExcludeGenerated(v bool) func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeeds) WithPretty() func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeeds) WithHuman() func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeeds) WithErrorTrace() func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeeds) WithFilterPath(v ...string) func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeeds) WithHeader(h map[string]string) func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeeds) WithOpaqueID(s string) func(*MLGetDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}
