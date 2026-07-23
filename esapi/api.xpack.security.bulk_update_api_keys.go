package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityBulkUpdateAPIKeysFunc(t Transport) SecurityBulkUpdateAPIKeys {
	_ = "STUB: not implemented"
	return *new(SecurityBulkUpdateAPIKeys)
}

type SecurityBulkUpdateAPIKeys func(body io.Reader, o ...func(*SecurityBulkUpdateAPIKeysRequest)) (*Response, error)

type SecurityBulkUpdateAPIKeysRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityBulkUpdateAPIKeysRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityBulkUpdateAPIKeys) WithContext(v context.Context) func(*SecurityBulkUpdateAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkUpdateAPIKeys) WithPretty() func(*SecurityBulkUpdateAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkUpdateAPIKeys) WithHuman() func(*SecurityBulkUpdateAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkUpdateAPIKeys) WithErrorTrace() func(*SecurityBulkUpdateAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkUpdateAPIKeys) WithFilterPath(v ...string) func(*SecurityBulkUpdateAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkUpdateAPIKeys) WithHeader(h map[string]string) func(*SecurityBulkUpdateAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityBulkUpdateAPIKeys) WithOpaqueID(s string) func(*SecurityBulkUpdateAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}
