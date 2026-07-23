package esapi

import (
	"context"
	"net/http"
)

func newSecurityDisableUserProfileFunc(t Transport) SecurityDisableUserProfile {
	_ = "STUB: not implemented"
	return *new(SecurityDisableUserProfile)
}

type SecurityDisableUserProfile func(uid string, o ...func(*SecurityDisableUserProfileRequest)) (*Response, error)

type SecurityDisableUserProfileRequest struct {
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

func (r SecurityDisableUserProfileRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityDisableUserProfile) WithContext(v context.Context) func(*SecurityDisableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUserProfile) WithRefresh(v string) func(*SecurityDisableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUserProfile) WithPretty() func(*SecurityDisableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUserProfile) WithHuman() func(*SecurityDisableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUserProfile) WithErrorTrace() func(*SecurityDisableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUserProfile) WithFilterPath(v ...string) func(*SecurityDisableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUserProfile) WithHeader(h map[string]string) func(*SecurityDisableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityDisableUserProfile) WithOpaqueID(s string) func(*SecurityDisableUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}
