package esapi

import (
	"context"
	"io"
	"net/http"
)

func newConnectorSecretPostFunc(t Transport) ConnectorSecretPost {
	_ = "STUB: not implemented"
	return *new(ConnectorSecretPost)
}

type ConnectorSecretPost func(body io.Reader, o ...func(*ConnectorSecretPostRequest)) (*Response, error)

type ConnectorSecretPostRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ConnectorSecretPostRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ConnectorSecretPost) WithContext(v context.Context) func(*ConnectorSecretPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPost) WithPretty() func(*ConnectorSecretPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPost) WithHuman() func(*ConnectorSecretPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPost) WithErrorTrace() func(*ConnectorSecretPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPost) WithFilterPath(v ...string) func(*ConnectorSecretPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPost) WithHeader(h map[string]string) func(*ConnectorSecretPostRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ConnectorSecretPost) WithOpaqueID(s string) func(*ConnectorSecretPostRequest) {
	_ = "STUB: not implemented"
	return nil
}
