package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesPutSettingsFunc(t Transport) IndicesPutSettings {
	_ = "STUB: not implemented"
	return *new(IndicesPutSettings)
}

type IndicesPutSettings func(body io.Reader, o ...func(*IndicesPutSettingsRequest)) (*Response, error)

type IndicesPutSettingsRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices    *bool
	ExpandWildcards   []string
	FlatSettings      *bool
	IgnoreUnavailable *bool
	MasterTimeout     time.Duration
	PreserveExisting  *bool
	Reopen            *bool
	Timeout           time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesPutSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPutSettings) WithContext(v context.Context) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithIndex(v ...string) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithAllowNoIndices(v bool) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithExpandWildcards(v ...string) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithFlatSettings(v bool) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithIgnoreUnavailable(v bool) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithMasterTimeout(v time.Duration) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithPreserveExisting(v bool) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithReopen(v bool) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithTimeout(v time.Duration) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithPretty() func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithHuman() func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithErrorTrace() func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithFilterPath(v ...string) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithHeader(h map[string]string) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutSettings) WithOpaqueID(s string) func(*IndicesPutSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
