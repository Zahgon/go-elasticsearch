package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateServiceDocumentTypeFunc(t Transport) ConnectorUpdateServiceDocumentType {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateServiceDocumentType)
}

type ConnectorUpdateServiceDocumentType func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateServiceDocumentTypeRequest)) (*Response, error)

type ConnectorUpdateServiceDocumentTypeRequest struct {
	Body io.Reader

	ConnectorID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorUpdateServiceDocumentTypeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateServiceDocumentType) WithContext(v context.Context) func(*ConnectorUpdateServiceDocumentTypeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateServiceDocumentType) WithPretty() func(*ConnectorUpdateServiceDocumentTypeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateServiceDocumentType) WithHuman() func(*ConnectorUpdateServiceDocumentTypeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateServiceDocumentType) WithErrorTrace() func(*ConnectorUpdateServiceDocumentTypeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateServiceDocumentType) WithFilterPath(v ...string) func(*ConnectorUpdateServiceDocumentTypeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateServiceDocumentType) WithHeader(h map[string]string) func(*ConnectorUpdateServiceDocumentTypeRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateServiceDocumentType) WithOpaqueID(s string) func(*ConnectorUpdateServiceDocumentTypeRequest) {
	_ = "STUB: not implemented"
	return nil
}
