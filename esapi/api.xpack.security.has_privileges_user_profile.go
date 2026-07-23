package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityHasPrivilegesUserProfileFunc(t Transport) SecurityHasPrivilegesUserProfile {
	_ = "STUB: not implemented"
	return *new(SecurityHasPrivilegesUserProfile)
}

type SecurityHasPrivilegesUserProfile func(body io.Reader, o ...func(*SecurityHasPrivilegesUserProfileRequest)) (*Response, error)

type SecurityHasPrivilegesUserProfileRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityHasPrivilegesUserProfileRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityHasPrivilegesUserProfile) WithContext(v context.Context) func(*SecurityHasPrivilegesUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivilegesUserProfile) WithPretty() func(*SecurityHasPrivilegesUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivilegesUserProfile) WithHuman() func(*SecurityHasPrivilegesUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivilegesUserProfile) WithErrorTrace() func(*SecurityHasPrivilegesUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivilegesUserProfile) WithFilterPath(v ...string) func(*SecurityHasPrivilegesUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivilegesUserProfile) WithHeader(h map[string]string) func(*SecurityHasPrivilegesUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityHasPrivilegesUserProfile) WithOpaqueID(s string) func(*SecurityHasPrivilegesUserProfileRequest) {
	_ = "STUB: not implemented"
	return nil
}
