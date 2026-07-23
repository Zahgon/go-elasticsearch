package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSecurityQueryAPIKeysFunc(t Transport) SecurityQueryAPIKeys {
	_ = "STUB: not implemented"
	return *new(SecurityQueryAPIKeys)
}

type SecurityQueryAPIKeys func(o ...func(*SecurityQueryAPIKeysRequest)) (*Response, error)

type SecurityQueryAPIKeysRequest struct {
	Body io.Reader

	TypedKeys      *bool
	WithLimitedBy  *bool
	WithProfileUID *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityQueryAPIKeysRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityQueryAPIKeys) WithContext(v context.Context) func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithBody(v io.Reader) func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithTypedKeys(v bool) func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithWithLimitedBy(v bool) func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithWithProfileUID(v bool) func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithPretty() func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithHuman() func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithErrorTrace() func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithFilterPath(v ...string) func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithHeader(h map[string]string) func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityQueryAPIKeys) WithOpaqueID(s string) func(*SecurityQueryAPIKeysRequest) {
	_ = "STUB: not implemented"
	return nil
}
