package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetUserProfileFunc(t Transport) SecurityGetUserProfile {
	_ = "STUB: not implemented"
	return *new(SecurityGetUserProfile)
}

type SecurityGetUserProfile func(uid []string, o ...func(*SecurityGetUserProfileRequest)) (*Response, error)

type SecurityGetUserProfileRequest struct {
	UID []string

	Data []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetUserProfileRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetUserProfile) WithContext(v context.Context) func(*SecurityGetUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserProfile) WithData(v ...string) func(*SecurityGetUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserProfile) WithPretty() func(*SecurityGetUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserProfile) WithHuman() func(*SecurityGetUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserProfile) WithErrorTrace() func(*SecurityGetUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserProfile) WithFilterPath(v ...string) func(*SecurityGetUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserProfile) WithHeader(h map[string]string) func(*SecurityGetUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetUserProfile) WithOpaqueID(s string) func(*SecurityGetUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}
