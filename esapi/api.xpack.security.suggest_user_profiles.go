package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecuritySuggestUserProfilesFunc(t Transport) SecuritySuggestUserProfiles {
	_ = "STUB: not implemented"
	return *new(SecuritySuggestUserProfiles)
}

type SecuritySuggestUserProfiles func(o ...func(*SecuritySuggestUserProfilesRequest)) (*Response, error)

type SecuritySuggestUserProfilesRequest struct {
	Body io.Reader

	Data []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecuritySuggestUserProfilesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecuritySuggestUserProfiles) WithContext(v context.Context) func(*SecuritySuggestUserProfilesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySuggestUserProfiles) WithBody(v io.Reader) func(*SecuritySuggestUserProfilesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySuggestUserProfiles) WithData(v ...string) func(*SecuritySuggestUserProfilesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySuggestUserProfiles) WithPretty() func(*SecuritySuggestUserProfilesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySuggestUserProfiles) WithHuman() func(*SecuritySuggestUserProfilesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySuggestUserProfiles) WithErrorTrace() func(*SecuritySuggestUserProfilesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySuggestUserProfiles) WithFilterPath(v ...string) func(*SecuritySuggestUserProfilesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySuggestUserProfiles) WithHeader(h map[string]string) func(*SecuritySuggestUserProfilesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySuggestUserProfiles) WithOpaqueID(s string) func(*SecuritySuggestUserProfilesRequest) {
	_ = "STUB: not implemented"
	return nil
}
