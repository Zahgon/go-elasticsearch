package esapi

import (
	"context"
	"io"
	"net/http"
)

func newProjectTagsFunc(t Transport) ProjectTags {
	_ = "STUB: not implemented"
	return *new(ProjectTags)
}

type ProjectTags func(o ...func(*ProjectTagsRequest)) (*Response, error)

type ProjectTagsRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProjectTagsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProjectTags) WithContext(v context.Context) func(*ProjectTagsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectTags) WithBody(v io.Reader) func(*ProjectTagsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectTags) WithPretty() func(*ProjectTagsRequest) { _ = "STUB: not implemented"; return nil }

func (f ProjectTags) WithHuman() func(*ProjectTagsRequest) { _ = "STUB: not implemented"; return nil }

func (f ProjectTags) WithErrorTrace() func(*ProjectTagsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectTags) WithFilterPath(v ...string) func(*ProjectTagsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectTags) WithHeader(h map[string]string) func(*ProjectTagsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectTags) WithOpaqueID(s string) func(*ProjectTagsRequest) {
	_ = "STUB: not implemented"
	return nil
}
