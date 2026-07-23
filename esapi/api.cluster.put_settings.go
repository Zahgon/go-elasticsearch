package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newClusterPutSettingsFunc(t Transport) ClusterPutSettings {
	_ = "STUB: not implemented"
	return *new(ClusterPutSettings)
}

type ClusterPutSettings func(body io.Reader, o ...func(*ClusterPutSettingsRequest)) (*Response, error)

type ClusterPutSettingsRequest struct {
	Body io.Reader

	FlatSettings  *bool
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

func (r ClusterPutSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterPutSettings) WithContext(v context.Context) func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutSettings) WithFlatSettings(v bool) func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutSettings) WithMasterTimeout(v time.Duration) func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutSettings) WithTimeout(v time.Duration) func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutSettings) WithPretty() func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutSettings) WithHuman() func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutSettings) WithErrorTrace() func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutSettings) WithFilterPath(v ...string) func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutSettings) WithHeader(h map[string]string) func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutSettings) WithOpaqueID(s string) func(*ClusterPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
