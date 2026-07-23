package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesPutDataStreamSettingsFunc(t Transport) IndicesPutDataStreamSettings {
	_ = "STUB: not implemented"
	return *new(IndicesPutDataStreamSettings)
}

type IndicesPutDataStreamSettings func(name []string, body io.Reader, o ...func(*IndicesPutDataStreamSettingsRequest)) (*Response, error)

type IndicesPutDataStreamSettingsRequest struct {
	Body io.Reader

	Name []string

	DryRun        *bool
	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesPutDataStreamSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPutDataStreamSettings) WithContext(v context.Context) func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamSettings) WithDryRun(v bool) func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamSettings) WithMasterTimeout(v time.Duration) func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamSettings) WithTimeout(v time.Duration) func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamSettings) WithPretty() func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamSettings) WithHuman() func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamSettings) WithErrorTrace() func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamSettings) WithFilterPath(v ...string) func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamSettings) WithHeader(h map[string]string) func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamSettings) WithOpaqueID(s string) func(*IndicesPutDataStreamSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
