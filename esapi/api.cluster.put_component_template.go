package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newClusterPutComponentTemplateFunc(t Transport) ClusterPutComponentTemplate {
	_ = "STUB: not implemented"
	return *new(ClusterPutComponentTemplate)
}

type ClusterPutComponentTemplate func(name string, body io.Reader, o ...func(*ClusterPutComponentTemplateRequest)) (*Response, error)

type ClusterPutComponentTemplateRequest struct {
	Body io.Reader

	Name string

	Cause         string
	Create        *bool
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterPutComponentTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterPutComponentTemplate) WithContext(v context.Context) func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutComponentTemplate) WithCause(v string) func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutComponentTemplate) WithCreate(v bool) func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutComponentTemplate) WithMasterTimeout(v time.Duration) func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutComponentTemplate) WithPretty() func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutComponentTemplate) WithHuman() func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutComponentTemplate) WithErrorTrace() func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutComponentTemplate) WithFilterPath(v ...string) func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutComponentTemplate) WithHeader(h map[string]string) func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterPutComponentTemplate) WithOpaqueID(s string) func(*ClusterPutComponentTemplateRequest) {
	_ = "STUB: not implemented"
	return nil
}
