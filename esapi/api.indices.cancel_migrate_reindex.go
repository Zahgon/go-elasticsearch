package esapi

import (
	"context"
	"net/http"
)

func newIndicesCancelMigrateReindexFunc(t Transport) IndicesCancelMigrateReindex {
	_ = "STUB: not implemented"
	return *new(IndicesCancelMigrateReindex)
}

type IndicesCancelMigrateReindex func(index []string, o ...func(*IndicesCancelMigrateReindexRequest)) (*Response, error)

type IndicesCancelMigrateReindexRequest struct {
	Index []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesCancelMigrateReindexRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesCancelMigrateReindex) WithContext(v context.Context) func(*IndicesCancelMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCancelMigrateReindex) WithPretty() func(*IndicesCancelMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCancelMigrateReindex) WithHuman() func(*IndicesCancelMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCancelMigrateReindex) WithErrorTrace() func(*IndicesCancelMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCancelMigrateReindex) WithFilterPath(v ...string) func(*IndicesCancelMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCancelMigrateReindex) WithHeader(h map[string]string) func(*IndicesCancelMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCancelMigrateReindex) WithOpaqueID(s string) func(*IndicesCancelMigrateReindexRequest) {
	_ = "STUB: not implemented"
	return nil
}
