package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTransformUpgradeTransformsFunc(t Transport) TransformUpgradeTransforms {
	_ = "STUB: not implemented"
	return *new(TransformUpgradeTransforms)
}

type TransformUpgradeTransforms func(o ...func(*TransformUpgradeTransformsRequest)) (*Response, error)

type TransformUpgradeTransformsRequest struct {
	DryRun  *bool
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformUpgradeTransformsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformUpgradeTransforms) WithContext(v context.Context) func(*TransformUpgradeTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpgradeTransforms) WithDryRun(v bool) func(*TransformUpgradeTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpgradeTransforms) WithTimeout(v time.Duration) func(*TransformUpgradeTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpgradeTransforms) WithPretty() func(*TransformUpgradeTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpgradeTransforms) WithHuman() func(*TransformUpgradeTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpgradeTransforms) WithErrorTrace() func(*TransformUpgradeTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpgradeTransforms) WithFilterPath(v ...string) func(*TransformUpgradeTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpgradeTransforms) WithHeader(h map[string]string) func(*TransformUpgradeTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformUpgradeTransforms) WithOpaqueID(s string) func(*TransformUpgradeTransformsRequest) {
	_ = "STUB: not implemented"
	return nil
}
