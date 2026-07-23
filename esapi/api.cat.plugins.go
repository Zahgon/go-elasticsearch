package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatPluginsFunc(t Transport) CatPlugins { _ = "STUB: not implemented"; return *new(CatPlugins) }

type CatPlugins func(o ...func(*CatPluginsRequest)) (*Response, error)

type CatPluginsRequest struct {
	Bytes            string
	Format           string
	H                []string
	Help             *bool
	IncludeBootstrap *bool
	Local            *bool
	MasterTimeout    time.Duration
	S                []string
	Time             string
	V                *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatPluginsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatPlugins) WithContext(v context.Context) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithBytes(v string) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithFormat(v string) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithH(v ...string) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithHelp(v bool) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithIncludeBootstrap(v bool) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithLocal(v bool) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithMasterTimeout(v time.Duration) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithS(v ...string) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithTime(v string) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithV(v bool) func(*CatPluginsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatPlugins) WithPretty() func(*CatPluginsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatPlugins) WithHuman() func(*CatPluginsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatPlugins) WithErrorTrace() func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithFilterPath(v ...string) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithHeader(h map[string]string) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatPlugins) WithOpaqueID(s string) func(*CatPluginsRequest) {
	_ = "STUB: not implemented"
	return nil
}
