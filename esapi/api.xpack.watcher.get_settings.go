package esapi

import (
	"context"
	"net/http"
	"time"
)

func newWatcherGetSettingsFunc(t Transport) WatcherGetSettings {
	_ = "STUB: not implemented"
	return *new(WatcherGetSettings)
}

type WatcherGetSettings func(o ...func(*WatcherGetSettingsRequest)) (*Response, error)

type WatcherGetSettingsRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherGetSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherGetSettings) WithContext(v context.Context) func(*WatcherGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetSettings) WithMasterTimeout(v time.Duration) func(*WatcherGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetSettings) WithPretty() func(*WatcherGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetSettings) WithHuman() func(*WatcherGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetSettings) WithErrorTrace() func(*WatcherGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetSettings) WithFilterPath(v ...string) func(*WatcherGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetSettings) WithHeader(h map[string]string) func(*WatcherGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetSettings) WithOpaqueID(s string) func(*WatcherGetSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
