package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newWatcherUpdateSettingsFunc(t Transport) WatcherUpdateSettings {
	_ = "STUB: not implemented"
	return *new(WatcherUpdateSettings)
}

type WatcherUpdateSettings func(body io.Reader, o ...func(*WatcherUpdateSettingsRequest)) (*Response, error)

type WatcherUpdateSettingsRequest struct {
	Body io.Reader

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherUpdateSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherUpdateSettings) WithContext(v context.Context) func(*WatcherUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherUpdateSettings) WithMasterTimeout(v time.Duration) func(*WatcherUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherUpdateSettings) WithTimeout(v time.Duration) func(*WatcherUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherUpdateSettings) WithPretty() func(*WatcherUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherUpdateSettings) WithHuman() func(*WatcherUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherUpdateSettings) WithErrorTrace() func(*WatcherUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherUpdateSettings) WithFilterPath(v ...string) func(*WatcherUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherUpdateSettings) WithHeader(h map[string]string) func(*WatcherUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherUpdateSettings) WithOpaqueID(s string) func(*WatcherUpdateSettingsRequest) {
	_ = "STUB: not implemented"
	return nil
}
