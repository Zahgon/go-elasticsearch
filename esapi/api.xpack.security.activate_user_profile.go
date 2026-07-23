package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityActivateUserProfileFunc(t Transport) SecurityActivateUserProfile {
	_ = "STUB: not implemented"
	return *new(SecurityActivateUserProfile)
}

type SecurityActivateUserProfile func(body io.Reader, o ...func(*SecurityActivateUserProfileRequest)) (*Response, error)

type SecurityActivateUserProfileRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityActivateUserProfileRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityActivateUserProfile) WithContext(v context.Context) func(*SecurityActivateUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityActivateUserProfile) WithPretty() func(*SecurityActivateUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityActivateUserProfile) WithHuman() func(*SecurityActivateUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityActivateUserProfile) WithErrorTrace() func(*SecurityActivateUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityActivateUserProfile) WithFilterPath(v ...string) func(*SecurityActivateUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityActivateUserProfile) WithHeader(h map[string]string) func(*SecurityActivateUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityActivateUserProfile) WithOpaqueID(s string) func(*SecurityActivateUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}
