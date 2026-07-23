package esapi

import (
	"context"
	"net/http"
)

func newMLPutTrainedModelAliasFunc(t Transport) MLPutTrainedModelAlias {
	_ = "STUB: not implemented"
	return *new(MLPutTrainedModelAlias)
}

type MLPutTrainedModelAlias func(model_alias string, model_id string, o ...func(*MLPutTrainedModelAliasRequest)) (*Response, error)

type MLPutTrainedModelAliasRequest struct {
	ModelAlias string
	ModelID    string

	Reassign *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLPutTrainedModelAliasRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutTrainedModelAlias) WithContext(v context.Context) func(*MLPutTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelAlias) WithReassign(v bool) func(*MLPutTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelAlias) WithPretty() func(*MLPutTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelAlias) WithHuman() func(*MLPutTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelAlias) WithErrorTrace() func(*MLPutTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelAlias) WithFilterPath(v ...string) func(*MLPutTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelAlias) WithHeader(h map[string]string) func(*MLPutTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelAlias) WithOpaqueID(s string) func(*MLPutTrainedModelAliasRequest) {
	_ = "STUB: not implemented"
	return nil
}
