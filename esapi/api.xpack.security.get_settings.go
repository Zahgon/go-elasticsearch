package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSecurityGetSettingsFunc(t Transport) SecurityGetSettings {
	_ = "STUB: not implemented"
	return *new(SecurityGetSettings)
}

type SecurityGetSettings func(o ...func(*SecurityGetSettingsRequest)) (*Response, error)

type SecurityGetSettingsRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetSettings) WithContext(v context.Context) func(*SecurityGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetSettings) WithMasterTimeout(v time.Duration) func(*SecurityGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetSettings) WithPretty() func(*SecurityGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetSettings) WithHuman() func(*SecurityGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetSettings) WithErrorTrace() func(*SecurityGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetSettings) WithFilterPath(v ...string) func(*SecurityGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetSettings) WithHeader(h map[string]string) func(*SecurityGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetSettings) WithOpaqueID(s string) func(*SecurityGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
