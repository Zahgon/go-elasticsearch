package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPostDataFunc(t Transport) MLPostData { _ = "STUB: not implemented"; return *new(MLPostData) }

type MLPostData func(job_id string, body io.Reader, o ...func(*MLPostDataRequest)) (*Response, error)

type MLPostDataRequest struct {
	Body io.Reader

	JobID string

	ResetEnd   string
	ResetStart string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLPostDataRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPostData) WithContext(v context.Context) func(*MLPostDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostData) WithResetEnd(v string) func(*MLPostDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostData) WithResetStart(v string) func(*MLPostDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostData) WithPretty() func(*MLPostDataRequest) { _ = "STUB: not implemented"; return nil }

func (f MLPostData) WithHuman() func(*MLPostDataRequest) { _ = "STUB: not implemented"; return nil }

func (f MLPostData) WithErrorTrace() func(*MLPostDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostData) WithFilterPath(v ...string) func(*MLPostDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostData) WithHeader(h map[string]string) func(*MLPostDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostData) WithOpaqueID(s string) func(*MLPostDataRequest) {
	_ = "STUB: not implemented"
	return nil
}
