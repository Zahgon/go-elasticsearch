package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLValidateDetectorFunc(t Transport) MLValidateDetector {
	_ = "STUB: not implemented"
	return *new(MLValidateDetector)
}

type MLValidateDetector func(body io.Reader, o ...func(*MLValidateDetectorRequest)) (*Response, error)

type MLValidateDetectorRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLValidateDetectorRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLValidateDetector) WithContext(v context.Context) func(*MLValidateDetectorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidateDetector) WithPretty() func(*MLValidateDetectorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidateDetector) WithHuman() func(*MLValidateDetectorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidateDetector) WithErrorTrace() func(*MLValidateDetectorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidateDetector) WithFilterPath(v ...string) func(*MLValidateDetectorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidateDetector) WithHeader(h map[string]string) func(*MLValidateDetectorRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLValidateDetector) WithOpaqueID(s string) func(*MLValidateDetectorRequest) {
	_ = "STUB: not implemented"
	return nil
}
