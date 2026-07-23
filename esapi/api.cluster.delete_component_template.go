package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterDeleteComponentTemplateFunc(t Transport) ClusterDeleteComponentTemplate {
	_ = "STUB: not implemented"
	return *new(ClusterDeleteComponentTemplate)
}

type ClusterDeleteComponentTemplate func(name []string, o ...func(*ClusterDeleteComponentTemplateRequest)) (*Response, error)

type ClusterDeleteComponentTemplateRequest struct {
	Name []string

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

func (r ClusterDeleteComponentTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterDeleteComponentTemplate) WithContext(v context.Context) func(*ClusterDeleteComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteComponentTemplate) WithMasterTimeout(v time.Duration) func(*ClusterDeleteComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteComponentTemplate) WithTimeout(v time.Duration) func(*ClusterDeleteComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteComponentTemplate) WithPretty() func(*ClusterDeleteComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteComponentTemplate) WithHuman() func(*ClusterDeleteComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteComponentTemplate) WithErrorTrace() func(*ClusterDeleteComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteComponentTemplate) WithFilterPath(v ...string) func(*ClusterDeleteComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteComponentTemplate) WithHeader(h map[string]string) func(*ClusterDeleteComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterDeleteComponentTemplate) WithOpaqueID(s string) func(*ClusterDeleteComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
