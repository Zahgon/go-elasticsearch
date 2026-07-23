package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetDataStreamSettingsFunc(t Transport) IndicesGetDataStreamSettings {
	_ = "STUB: not implemented"
	return *new(IndicesGetDataStreamSettings)
}

type IndicesGetDataStreamSettings func(name []string, o ...func(*IndicesGetDataStreamSettingsRequest)) (*Response, error)

type IndicesGetDataStreamSettingsRequest struct {
	Name []string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetDataStreamSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetDataStreamSettings) WithContext(v context.Context) func(*IndicesGetDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamSettings) WithMasterTimeout(v time.Duration) func(*IndicesGetDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamSettings) WithPretty() func(*IndicesGetDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamSettings) WithHuman() func(*IndicesGetDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamSettings) WithErrorTrace() func(*IndicesGetDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamSettings) WithFilterPath(v ...string) func(*IndicesGetDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamSettings) WithHeader(h map[string]string) func(*IndicesGetDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamSettings) WithOpaqueID(s string) func(*IndicesGetDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
