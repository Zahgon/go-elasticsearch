package esapi

import (
	"context"
	"net/http"
	"time"
)

func newFeaturesGetFeaturesFunc(t Transport) FeaturesGetFeatures {
	_ = "STUB: not implemented"
	return *new(FeaturesGetFeatures)
}

type FeaturesGetFeatures func(o ...func(*FeaturesGetFeaturesRequest)) (*Response, error)

type FeaturesGetFeaturesRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r FeaturesGetFeaturesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FeaturesGetFeatures) WithContext(v context.Context) func(*FeaturesGetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesGetFeatures) WithMasterTimeout(v time.Duration) func(*FeaturesGetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesGetFeatures) WithPretty() func(*FeaturesGetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesGetFeatures) WithHuman() func(*FeaturesGetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesGetFeatures) WithErrorTrace() func(*FeaturesGetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesGetFeatures) WithFilterPath(v ...string) func(*FeaturesGetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesGetFeatures) WithHeader(h map[string]string) func(*FeaturesGetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FeaturesGetFeatures) WithOpaqueID(s string) func(*FeaturesGetFeaturesRequest) {
	_ = "STUB: not implemented"
	return nil
}
