package esapi

import (
	"context"
	"net/http"
)

func newMLGetTrainedModelsStatsFunc(t Transport) MLGetTrainedModelsStats {
	_ = "STUB: not implemented"
	return *new(MLGetTrainedModelsStats)
}

type MLGetTrainedModelsStats func(o ...func(*MLGetTrainedModelsStatsRequest)) (*Response, error)

type MLGetTrainedModelsStatsRequest struct {
	ModelID []string

	AllowNoMatch *bool
	From         *int
	Size         *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetTrainedModelsStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetTrainedModelsStats) WithContext(v context.Context) func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithModelID(v ...string) func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithAllowNoMatch(v bool) func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithFrom(v int) func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithSize(v int) func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithPretty() func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithHuman() func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithErrorTrace() func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithFilterPath(v ...string) func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithHeader(h map[string]string) func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetTrainedModelsStats) WithOpaqueID(s string) func(*MLGetTrainedModelsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
