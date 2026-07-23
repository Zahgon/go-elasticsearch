package esapi

import (
	"context"
	"net/http"
	"time"
)

func newFeaturesResetFeaturesFunc(t Transport) FeaturesResetFeatures {
	_ = "STUB: not implemented"
	return *new(FeaturesResetFeatures)
}

type FeaturesResetFeatures func(o ...func(*FeaturesResetFeaturesRequest)) (*Response, error)

type FeaturesResetFeaturesRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r FeaturesResetFeaturesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FeaturesResetFeatures) WithContext(v context.Context) func(*FeaturesResetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesResetFeatures) WithMasterTimeout(v time.Duration) func(*FeaturesResetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesResetFeatures) WithPretty() func(*FeaturesResetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesResetFeatures) WithHuman() func(*FeaturesResetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesResetFeatures) WithErrorTrace() func(*FeaturesResetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesResetFeatures) WithFilterPath(v ...string) func(*FeaturesResetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesResetFeatures) WithHeader(h map[string]string) func(*FeaturesResetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesResetFeatures) WithOpaqueID(s string) func(*FeaturesResetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}
