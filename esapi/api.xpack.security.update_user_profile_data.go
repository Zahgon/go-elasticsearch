package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityUpdateUserProfileDataFunc(t Transport) SecurityUpdateUserProfileData {
	_ = "STUB: not implemented"
	return *new(SecurityUpdateUserProfileData)
}

type SecurityUpdateUserProfileData func(body io.Reader, uid string, o ...func(*SecurityUpdateUserProfileDataRequest)) (*Response, error)

type SecurityUpdateUserProfileDataRequest struct {
	Body io.Reader

	UID string

	IfPrimaryTerm *int64
	IfSeqNo       *int64
	Refresh       string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityUpdateUserProfileDataRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityUpdateUserProfileData) WithContext(v context.Context) func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateUserProfileData) WithIfPrimaryTerm(v int64) func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateUserProfileData) WithIfSeqNo(v int64) func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateUserProfileData) WithRefresh(v string) func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateUserProfileData) WithPretty() func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateUserProfileData) WithHuman() func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateUserProfileData) WithErrorTrace() func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateUserProfileData) WithFilterPath(v ...string) func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateUserProfileData) WithHeader(h map[string]string) func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateUserProfileData) WithOpaqueID(s string) func(*SecurityUpdateUserProfileDataRequest) {
	_ = "STUB: not implemented"
	return nil
}
