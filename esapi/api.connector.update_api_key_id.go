package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorUpdateAPIKeyDocumentIDFunc(t Transport) ConnectorUpdateAPIKeyDocumentID {
	_ = "STUB: not implemented"
	return *new(ConnectorUpdateAPIKeyDocumentID)
}

type ConnectorUpdateAPIKeyDocumentID func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateAPIKeyDocumentIDRequest)) (*Response, error)

type ConnectorUpdateAPIKeyDocumentIDRequest struct {
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

func (r ConnectorUpdateAPIKeyDocumentIDRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorUpdateAPIKeyDocumentID) WithContext(v context.Context) func(*ConnectorUpdateAPIKeyDocumentIDRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateAPIKeyDocumentID) WithPretty() func(*ConnectorUpdateAPIKeyDocumentIDRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateAPIKeyDocumentID) WithHuman() func(*ConnectorUpdateAPIKeyDocumentIDRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateAPIKeyDocumentID) WithErrorTrace() func(*ConnectorUpdateAPIKeyDocumentIDRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateAPIKeyDocumentID) WithFilterPath(v ...string) func(*ConnectorUpdateAPIKeyDocumentIDRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateAPIKeyDocumentID) WithHeader(h map[string]string) func(*ConnectorUpdateAPIKeyDocumentIDRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorUpdateAPIKeyDocumentID) WithOpaqueID(s string) func(*ConnectorUpdateAPIKeyDocumentIDRequest) {
	_ = "STUB: not implemented"
	return nil
}
