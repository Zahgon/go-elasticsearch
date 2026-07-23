package esapi

import (
	"context"
	"net/http"
)

func newSecuritySamlServiceProviderMetadataFunc(t Transport) SecuritySamlServiceProviderMetadata {
	_ = "STUB: not implemented"
	return *new(SecuritySamlServiceProviderMetadata)
}

type SecuritySamlServiceProviderMetadata func(realm_name string, o ...func(*SecuritySamlServiceProviderMetadataRequest)) (*Response, error)

type SecuritySamlServiceProviderMetadataRequest struct {
	RealmName string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecuritySamlServiceProviderMetadataRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecuritySamlServiceProviderMetadata) WithContext(v context.Context) func(*SecuritySamlServiceProviderMetadataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlServiceProviderMetadata) WithPretty() func(*SecuritySamlServiceProviderMetadataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlServiceProviderMetadata) WithHuman() func(*SecuritySamlServiceProviderMetadataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlServiceProviderMetadata) WithErrorTrace() func(*SecuritySamlServiceProviderMetadataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlServiceProviderMetadata) WithFilterPath(v ...string) func(*SecuritySamlServiceProviderMetadataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlServiceProviderMetadata) WithHeader(h map[string]string) func(*SecuritySamlServiceProviderMetadataRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecuritySamlServiceProviderMetadata) WithOpaqueID(s string) func(*SecuritySamlServiceProviderMetadataRequest) {
	_ = "STUB: not implemented"
	return nil
}
