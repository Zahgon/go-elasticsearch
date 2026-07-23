package esapi

import (
	"context"
	"net/http"
	"time"
)

func newMLSetUpgradeModeFunc(t Transport) MLSetUpgradeMode {
	_ = "STUB: not implemented"
	return *new(MLSetUpgradeMode)
}

type MLSetUpgradeMode func(o ...func(*MLSetUpgradeModeRequest)) (*Response, error)

type MLSetUpgradeModeRequest struct {
	Enabled *bool
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLSetUpgradeModeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLSetUpgradeMode) WithContext(v context.Context) func(*MLSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLSetUpgradeMode) WithEnabled(v bool) func(*MLSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLSetUpgradeMode) WithTimeout(v time.Duration) func(*MLSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLSetUpgradeMode) WithPretty() func(*MLSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLSetUpgradeMode) WithHuman() func(*MLSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLSetUpgradeMode) WithErrorTrace() func(*MLSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLSetUpgradeMode) WithFilterPath(v ...string) func(*MLSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLSetUpgradeMode) WithHeader(h map[string]string) func(*MLSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLSetUpgradeMode) WithOpaqueID(s string) func(*MLSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}
