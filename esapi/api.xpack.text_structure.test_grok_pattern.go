package esapi

import (
	"context"
	"io"
	"net/http"
)

func newTextStructureTestGrokPatternFunc(t Transport) TextStructureTestGrokPattern {
	_ = "STUB: not implemented"
	return *new(TextStructureTestGrokPattern)
}

type TextStructureTestGrokPattern func(body io.Reader, o ...func(*TextStructureTestGrokPatternRequest)) (*Response, error)

type TextStructureTestGrokPatternRequest struct {
	Body io.Reader

	EcsCompatibility string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TextStructureTestGrokPatternRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TextStructureTestGrokPattern) WithContext(v context.Context) func(*TextStructureTestGrokPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureTestGrokPattern) WithEcsCompatibility(v string) func(*TextStructureTestGrokPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureTestGrokPattern) WithPretty() func(*TextStructureTestGrokPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureTestGrokPattern) WithHuman() func(*TextStructureTestGrokPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureTestGrokPattern) WithErrorTrace() func(*TextStructureTestGrokPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureTestGrokPattern) WithFilterPath(v ...string) func(*TextStructureTestGrokPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureTestGrokPattern) WithHeader(h map[string]string) func(*TextStructureTestGrokPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TextStructureTestGrokPattern) WithOpaqueID(s string) func(*TextStructureTestGrokPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}
