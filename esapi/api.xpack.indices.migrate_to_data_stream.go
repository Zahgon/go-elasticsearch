package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesMigrateToDataStreamFunc(t Transport) IndicesMigrateToDataStream {
	_ = "STUB: not implemented"
	return *new(IndicesMigrateToDataStream)
}

type IndicesMigrateToDataStream func(name string, o ...func(*IndicesMigrateToDataStreamRequest)) (*Response, error)

type IndicesMigrateToDataStreamRequest struct {
	Name string

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

func (r IndicesMigrateToDataStreamRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesMigrateToDataStream) WithContext(v context.Context) func(*IndicesMigrateToDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateToDataStream) WithMasterTimeout(v time.Duration) func(*IndicesMigrateToDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateToDataStream) WithTimeout(v time.Duration) func(*IndicesMigrateToDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateToDataStream) WithPretty() func(*IndicesMigrateToDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateToDataStream) WithHuman() func(*IndicesMigrateToDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateToDataStream) WithErrorTrace() func(*IndicesMigrateToDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateToDataStream) WithFilterPath(v ...string) func(*IndicesMigrateToDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateToDataStream) WithHeader(h map[string]string) func(*IndicesMigrateToDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesMigrateToDataStream) WithOpaqueID(s string) func(*IndicesMigrateToDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}
