package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterGetSettingsFunc(t Transport) ClusterGetSettings {
	_ = "STUB: not implemented"
	return *new(ClusterGetSettings)
}

type ClusterGetSettings func(o ...func(*ClusterGetSettingsRequest)) (*Response, error)

type ClusterGetSettingsRequest struct {
	FlatSettings    *bool
	IncludeDefaults *bool
	MasterTimeout   time.Duration
	Timeout         time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterGetSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterGetSettings) WithContext(v context.Context) func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithFlatSettings(v bool) func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithIncludeDefaults(v bool) func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithMasterTimeout(v time.Duration) func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithTimeout(v time.Duration) func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithPretty() func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithHuman() func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithErrorTrace() func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithFilterPath(v ...string) func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithHeader(h map[string]string) func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetSettings) WithOpaqueID(s string) func(*ClusterGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
