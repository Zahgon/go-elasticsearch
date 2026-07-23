package esapi

import (
	"context"
	"net/http"
)

func newSecurityClearCachedRealmsFunc(t Transport) SecurityClearCachedRealms {
	_ = "STUB: not implemented"
	return *new(SecurityClearCachedRealms)
}

type SecurityClearCachedRealms func(realms []string, o ...func(*SecurityClearCachedRealmsRequest)) (*Response, error)

type SecurityClearCachedRealmsRequest struct {
	Realms []string

	Usernames []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityClearCachedRealmsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityClearCachedRealms) WithContext(v context.Context) func(*SecurityClearCachedRealmsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRealms) WithUsernames(v ...string) func(*SecurityClearCachedRealmsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRealms) WithPretty() func(*SecurityClearCachedRealmsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRealms) WithHuman() func(*SecurityClearCachedRealmsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRealms) WithErrorTrace() func(*SecurityClearCachedRealmsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRealms) WithFilterPath(v ...string) func(*SecurityClearCachedRealmsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRealms) WithHeader(h map[string]string) func(*SecurityClearCachedRealmsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityClearCachedRealms) WithOpaqueID(s string) func(*SecurityClearCachedRealmsRequest) {
	_ = "STUB: not implemented"
	return nil
}
