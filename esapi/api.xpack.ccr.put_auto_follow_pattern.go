package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newCCRPutAutoFollowPatternFunc(t Transport) CCRPutAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(CCRPutAutoFollowPattern)
}

type CCRPutAutoFollowPattern func(name string, body io.Reader, o ...func(*CCRPutAutoFollowPatternRequest)) (*Response, error)

type CCRPutAutoFollowPatternRequest struct {
	Body io.Reader

	Name string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CCRPutAutoFollowPatternRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRPutAutoFollowPattern) WithContext(v context.Context) func(*CCRPutAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPutAutoFollowPattern) WithMasterTimeout(v time.Duration) func(*CCRPutAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPutAutoFollowPattern) WithPretty() func(*CCRPutAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPutAutoFollowPattern) WithHuman() func(*CCRPutAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPutAutoFollowPattern) WithErrorTrace() func(*CCRPutAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPutAutoFollowPattern) WithFilterPath(v ...string) func(*CCRPutAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPutAutoFollowPattern) WithHeader(h map[string]string) func(*CCRPutAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRPutAutoFollowPattern) WithOpaqueID(s string) func(*CCRPutAutoFollowPatternRequest) {
	_ = "STUB: not implemented"
	return nil
}
