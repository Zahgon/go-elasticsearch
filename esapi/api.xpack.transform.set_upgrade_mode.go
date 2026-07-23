package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTransformSetUpgradeModeFunc(t Transport) TransformSetUpgradeMode {
	_ = "STUB: not implemented"
	return *new(TransformSetUpgradeMode)
}

type TransformSetUpgradeMode func(o ...func(*TransformSetUpgradeModeRequest)) (*Response, error)

type TransformSetUpgradeModeRequest struct {
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

func (r TransformSetUpgradeModeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformSetUpgradeMode) WithContext(v context.Context) func(*TransformSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformSetUpgradeMode) WithEnabled(v bool) func(*TransformSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformSetUpgradeMode) WithTimeout(v time.Duration) func(*TransformSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformSetUpgradeMode) WithPretty() func(*TransformSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformSetUpgradeMode) WithHuman() func(*TransformSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformSetUpgradeMode) WithErrorTrace() func(*TransformSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformSetUpgradeMode) WithFilterPath(v ...string) func(*TransformSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformSetUpgradeMode) WithHeader(h map[string]string) func(*TransformSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformSetUpgradeMode) WithOpaqueID(s string) func(*TransformSetUpgradeModeRequest) {
	_ = "STUB: not implemented"
	return nil
}
