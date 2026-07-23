package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newSecurityUpdateSettingsFunc(t Transport) SecurityUpdateSettings {
	_ = "STUB: not implemented"
	return *new(SecurityUpdateSettings)
}

type SecurityUpdateSettings func(body io.Reader, o ...func(*SecurityUpdateSettingsRequest)) (*Response, error)

type SecurityUpdateSettingsRequest struct {
	Body io.Reader

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityUpdateSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityUpdateSettings) WithContext(v context.Context) func(*SecurityUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateSettings) WithMasterTimeout(v time.Duration) func(*SecurityUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateSettings) WithTimeout(v time.Duration) func(*SecurityUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateSettings) WithPretty() func(*SecurityUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateSettings) WithHuman() func(*SecurityUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateSettings) WithErrorTrace() func(*SecurityUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateSettings) WithFilterPath(v ...string) func(*SecurityUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateSettings) WithHeader(h map[string]string) func(*SecurityUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityUpdateSettings) WithOpaqueID(s string) func(*SecurityUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
