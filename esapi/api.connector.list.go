package esapi

import (
	"context"
	"net/http"
)

func newConnectorListFunc(t Transport) ConnectorList {
	_ = "STUB: not implemented"
	return *new(ConnectorList)
}

type ConnectorList func(o ...func(*ConnectorListRequest)) (*Response, error)

type ConnectorListRequest struct {
	ConnectorName  []string
	From           *int
	IncludeDeleted *bool
	IndexName      []string
	Query          string
	ServiceType    []string
	Size           *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorListRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorList) WithContext(v context.Context) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithConnectorName(v ...string) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithFrom(v int) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithIncludeDeleted(v bool) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithIndexName(v ...string) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithQuery(v string) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithServiceType(v ...string) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithSize(v int) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithPretty() func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithHuman() func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithErrorTrace() func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithFilterPath(v ...string) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithHeader(h map[string]string) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorList) WithOpaqueID(s string) func(*ConnectorListRequest) {
	_ = "STUB: not implemented"
	return nil
}
