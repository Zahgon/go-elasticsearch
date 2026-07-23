package esapi

import (
	"context"
	"net/http"
)

func newMLDeleteTrainedModelAliasFunc(t Transport) MLDeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return *new(MLDeleteTrainedModelAlias)
}

type MLDeleteTrainedModelAlias func(model_alias string, model_id string, o ...func(*MLDeleteTrainedModelAliasRequest)) (*Response, error)

type MLDeleteTrainedModelAliasRequest struct {
	ModelAlias string
	ModelID    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteTrainedModelAliasRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteTrainedModelAlias) WithContext(v context.Context) func(*MLDeleteTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModelAlias) WithPretty() func(*MLDeleteTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModelAlias) WithHuman() func(*MLDeleteTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModelAlias) WithErrorTrace() func(*MLDeleteTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModelAlias) WithFilterPath(v ...string) func(*MLDeleteTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModelAlias) WithHeader(h map[string]string) func(*MLDeleteTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModelAlias) WithOpaqueID(s string) func(*MLDeleteTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}
