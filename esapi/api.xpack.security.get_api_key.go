package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetAPIKeyFunc(t Transport) SecurityGetAPIKey {
	_ = "STUB: not implemented"
	return *new(SecurityGetAPIKey)
}

type SecurityGetAPIKey func(o ...func(*SecurityGetAPIKeyRequest)) (*Response, error)

type SecurityGetAPIKeyRequest struct {
	ActiveOnly     *bool
	ID             string
	Name           string
	Owner          *bool
	RealmName      string
	Username       string
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

func (r SecurityGetAPIKeyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetAPIKey) WithContext(v context.Context) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithActiveOnly(v bool) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithID(v string) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithName(v string) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithOwner(v bool) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithRealmName(v string) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithUsername(v string) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithWithLimitedBy(v bool) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithWithProfileUID(v bool) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithPretty() func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithHuman() func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithErrorTrace() func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithFilterPath(v ...string) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithHeader(h map[string]string) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetAPIKey) WithOpaqueID(s string) func(*SecurityGetAPIKeyRequest) {
	_ = "STUB: not implemented"
	return nil
}
