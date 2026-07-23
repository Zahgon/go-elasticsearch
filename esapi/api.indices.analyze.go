package esapi

import (
	"context"
	"io"
	"net/http"
)

func newIndicesAnalyzeFunc(t Transport) IndicesAnalyze {
	_ = "STUB: not implemented"
	return *new(IndicesAnalyze)
}

type IndicesAnalyze func(body io.Reader, o ...func(*IndicesAnalyzeRequest)) (*Response, error)

type IndicesAnalyzeRequest struct {
	Index string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesAnalyzeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesAnalyze) WithContext(v context.Context) func(*IndicesAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAnalyze) WithIndex(v string) func(*IndicesAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAnalyze) WithPretty() func(*IndicesAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAnalyze) WithHuman() func(*IndicesAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAnalyze) WithErrorTrace() func(*IndicesAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAnalyze) WithFilterPath(v ...string) func(*IndicesAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAnalyze) WithHeader(h map[string]string) func(*IndicesAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAnalyze) WithOpaqueID(s string) func(*IndicesAnalyzeRequest) {
	_ = "STUB: not implemented"
	return nil
}
