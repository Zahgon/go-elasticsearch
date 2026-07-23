package esapi

import (
	"context"
	"net/http"
)

func newSecurityClearAPIKeyCacheFunc(t Transport) SecurityClearAPIKeyCache {
	_ = "STUB: not implemented"
	return *new(SecurityClearAPIKeyCache)
}

type SecurityClearAPIKeyCache func(ids []string, o ...func(*SecurityClearAPIKeyCacheRequest)) (*Response, error)

type SecurityClearAPIKeyCacheRequest struct {
	Ids []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityClearAPIKeyCacheRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityClearAPIKeyCache) WithContext(v context.Context) func(*SecurityClearAPIKeyCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearAPIKeyCache) WithPretty() func(*SecurityClearAPIKeyCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearAPIKeyCache) WithHuman() func(*SecurityClearAPIKeyCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearAPIKeyCache) WithErrorTrace() func(*SecurityClearAPIKeyCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearAPIKeyCache) WithFilterPath(v ...string) func(*SecurityClearAPIKeyCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearAPIKeyCache) WithHeader(h map[string]string) func(*SecurityClearAPIKeyCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearAPIKeyCache) WithOpaqueID(s string) func(*SecurityClearAPIKeyCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}
