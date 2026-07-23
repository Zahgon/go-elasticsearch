package esapi

import (
	"context"
	"net/http"
)

func newSecurityEnableUserProfileFunc(t Transport) SecurityEnableUserProfile {
	_ = "STUB: not implemented"
	return *new(SecurityEnableUserProfile)
}

type SecurityEnableUserProfile func(uid string, o ...func(*SecurityEnableUserProfileRequest)) (*Response, error)

type SecurityEnableUserProfileRequest struct {
	UID string

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityEnableUserProfileRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityEnableUserProfile) WithContext(v context.Context) func(*SecurityEnableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUserProfile) WithRefresh(v string) func(*SecurityEnableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUserProfile) WithPretty() func(*SecurityEnableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUserProfile) WithHuman() func(*SecurityEnableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUserProfile) WithErrorTrace() func(*SecurityEnableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUserProfile) WithFilterPath(v ...string) func(*SecurityEnableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUserProfile) WithHeader(h map[string]string) func(*SecurityEnableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityEnableUserProfile) WithOpaqueID(s string) func(*SecurityEnableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}
