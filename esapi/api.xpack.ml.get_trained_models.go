package esapi

import (
	"context"
	"net/http"
)

func newMLGetTrainedModelsFunc(t Transport) MLGetTrainedModels {
	_ = "STUB: not implemented"
	return *new(MLGetTrainedModels)
}

type MLGetTrainedModels func(o ...func(*MLGetTrainedModelsRequest)) (*Response, error)

type MLGetTrainedModelsRequest struct {
	ModelID []string

	AllowNoMatch         *bool
	DecompressDefinition *bool
	ExcludeGenerated     *bool
	From                 *int
	Include              string
	Size                 *int
	Tags                 []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetTrainedModelsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetTrainedModels) WithContext(v context.Context) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithModelID(v ...string) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithAllowNoMatch(v bool) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithDecompressDefinition(v bool) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithExcludeGenerated(v bool) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithFrom(v int) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithInclude(v string) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithSize(v int) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithTags(v ...string) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithPretty() func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithHuman() func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithErrorTrace() func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithFilterPath(v ...string) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithHeader(h map[string]string) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModels) WithOpaqueID(s string) func(*MLGetTrainedModelsRequest) {
	_ = "STUB: not implemented"
	return nil
}
