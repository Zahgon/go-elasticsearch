package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPreviewDatafeedFunc(t Transport) MLPreviewDatafeed {
	_ = "STUB: not implemented"
	return *new(MLPreviewDatafeed)
}

type MLPreviewDatafeed func(o ...func(*MLPreviewDatafeedRequest)) (*Response, error)

type MLPreviewDatafeedRequest struct {
	Body io.Reader

	DatafeedID string

	End   string
	Start string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLPreviewDatafeedRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPreviewDatafeed) WithContext(v context.Context) func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithBody(v io.Reader) func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithDatafeedID(v string) func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithEnd(v string) func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithStart(v string) func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithPretty() func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithHuman() func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithErrorTrace() func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithFilterPath(v ...string) func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithHeader(h map[string]string) func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDatafeed) WithOpaqueID(s string) func(*MLPreviewDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}
