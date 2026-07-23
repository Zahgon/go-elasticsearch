package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLEvaluateDataFrameFunc(t Transport) MLEvaluateDataFrame {
	_ = "STUB: not implemented"
	return *new(MLEvaluateDataFrame)
}

type MLEvaluateDataFrame func(body io.Reader, o ...func(*MLEvaluateDataFrameRequest)) (*Response, error)

type MLEvaluateDataFrameRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLEvaluateDataFrameRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLEvaluateDataFrame) WithContext(v context.Context) func(*MLEvaluateDataFrameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEvaluateDataFrame) WithPretty() func(*MLEvaluateDataFrameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEvaluateDataFrame) WithHuman() func(*MLEvaluateDataFrameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEvaluateDataFrame) WithErrorTrace() func(*MLEvaluateDataFrameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEvaluateDataFrame) WithFilterPath(v ...string) func(*MLEvaluateDataFrameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEvaluateDataFrame) WithHeader(h map[string]string) func(*MLEvaluateDataFrameRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEvaluateDataFrame) WithOpaqueID(s string) func(*MLEvaluateDataFrameRequest) {
	_ = "STUB: not implemented"
	return nil
}
