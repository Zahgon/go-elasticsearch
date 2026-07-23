package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterGetComponentTemplateFunc(t Transport) ClusterGetComponentTemplate {
	_ = "STUB: not implemented"
	return *new(ClusterGetComponentTemplate)
}

type ClusterGetComponentTemplate func(o ...func(*ClusterGetComponentTemplateRequest)) (*Response, error)

type ClusterGetComponentTemplateRequest struct {
	Name string

	FlatSettings    *bool
	IncludeDefaults *bool
	Local           *bool
	MasterTimeout   time.Duration
	SettingsFilter  []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterGetComponentTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterGetComponentTemplate) WithContext(v context.Context) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithName(v string) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithFlatSettings(v bool) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithIncludeDefaults(v bool) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithLocal(v bool) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithMasterTimeout(v time.Duration) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithSettingsFilter(v ...string) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithPretty() func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithHuman() func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithErrorTrace() func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithFilterPath(v ...string) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithHeader(h map[string]string) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterGetComponentTemplate) WithOpaqueID(s string) func(*ClusterGetComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
