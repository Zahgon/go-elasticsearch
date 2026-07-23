package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetSettingsFunc(t Transport) IndicesGetSettings {
	_ = "STUB: not implemented"
	return *new(IndicesGetSettings)
}

type IndicesGetSettings func(o ...func(*IndicesGetSettingsRequest)) (*Response, error)

type IndicesGetSettingsRequest struct {
	Index []string

	Name []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	FlatSettings      *bool
	IgnoreUnavailable *bool
	IncludeDefaults   *bool
	Local             *bool
	MasterTimeout     time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetSettings) WithContext(v context.Context) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithIndex(v ...string) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithName(v ...string) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithAllowNoIndices(v bool) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithExpandWildcards(v ...string) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithFlatSettings(v bool) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithIgnoreUnavailable(v bool) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithIncludeDefaults(v bool) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithLocal(v bool) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithMasterTimeout(v time.Duration) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithPretty() func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithHuman() func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithErrorTrace() func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithFilterPath(v ...string) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithHeader(h map[string]string) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetSettings) WithOpaqueID(s string) func(*IndicesGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
