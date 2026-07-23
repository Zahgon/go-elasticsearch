package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterExistsComponentTemplateFunc(t Transport) ClusterExistsComponentTemplate {
	_ = "STUB: not implemented"
	return *new(ClusterExistsComponentTemplate)
}

type ClusterExistsComponentTemplate func(name []string, o ...func(*ClusterExistsComponentTemplateRequest)) (*Response, error)

type ClusterExistsComponentTemplateRequest struct {
	Name []string

	Local         *bool
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterExistsComponentTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterExistsComponentTemplate) WithContext(v context.Context) func(*ClusterExistsComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterExistsComponentTemplate) WithLocal(v bool) func(*ClusterExistsComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterExistsComponentTemplate) WithMasterTimeout(v time.Duration) func(*ClusterExistsComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterExistsComponentTemplate) WithPretty() func(*ClusterExistsComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterExistsComponentTemplate) WithHuman() func(*ClusterExistsComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterExistsComponentTemplate) WithErrorTrace() func(*ClusterExistsComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterExistsComponentTemplate) WithFilterPath(v ...string) func(*ClusterExistsComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterExistsComponentTemplate) WithHeader(h map[string]string) func(*ClusterExistsComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterExistsComponentTemplate) WithOpaqueID(s string) func(*ClusterExistsComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
