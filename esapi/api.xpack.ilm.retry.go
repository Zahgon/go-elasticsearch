package esapi

import (
	"context"
	"net/http"
)

func newILMRetryFunc(t Transport) ILMRetry { _ = "STUB: not implemented"; return *new(ILMRetry) }

type ILMRetry func(index string, o ...func(*ILMRetryRequest)) (*Response, error)

type ILMRetryRequest struct {
	Index string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ILMRetryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMRetry) WithContext(v context.Context) func(*ILMRetryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMRetry) WithPretty() func(*ILMRetryRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMRetry) WithHuman() func(*ILMRetryRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMRetry) WithErrorTrace() func(*ILMRetryRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMRetry) WithFilterPath(v ...string) func(*ILMRetryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMRetry) WithHeader(h map[string]string) func(*ILMRetryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMRetry) WithOpaqueID(s string) func(*ILMRetryRequest) {
	_ = "STUB: not implemented"
	return nil
}
