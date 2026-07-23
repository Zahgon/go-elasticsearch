package esapi

import (
	"context"
	"io"
	"net/http"
)

func newIndicesMigrateReindexFunc(t Transport) IndicesMigrateReindex {
	_ = "STUB: not implemented"
	return *new(IndicesMigrateReindex)
}

type IndicesMigrateReindex func(body io.Reader, o ...func(*IndicesMigrateReindexRequest)) (*Response, error)

type IndicesMigrateReindexRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesMigrateReindexRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesMigrateReindex) WithContext(v context.Context) func(*IndicesMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateReindex) WithPretty() func(*IndicesMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateReindex) WithHuman() func(*IndicesMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateReindex) WithErrorTrace() func(*IndicesMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateReindex) WithFilterPath(v ...string) func(*IndicesMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateReindex) WithHeader(h map[string]string) func(*IndicesMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateReindex) WithOpaqueID(s string) func(*IndicesMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}
