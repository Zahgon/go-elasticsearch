package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newScrollFunc(t Transport) Scroll { _ = "STUB: not implemented"; return *new(Scroll) }

type Scroll func(o ...func(*ScrollRequest)) (*Response, error)

type ScrollRequest struct {
	Body io.Reader

	ScrollID string

	RestTotalHitsAsInt *bool
	Scroll             time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ScrollRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Scroll) WithContext(v context.Context) func(*ScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Scroll) WithBody(v io.Reader) func(*ScrollRequest) { _ = "STUB: not implemented"; return nil }

func (f Scroll) WithScrollID(v string) func(*ScrollRequest) { _ = "STUB: not implemented"; return nil }

func (f Scroll) WithRestTotalHitsAsInt(v bool) func(*ScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Scroll) WithScroll(v time.Duration) func(*ScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Scroll) WithPretty() func(*ScrollRequest) { _ = "STUB: not implemented"; return nil }

func (f Scroll) WithHuman() func(*ScrollRequest) { _ = "STUB: not implemented"; return nil }

func (f Scroll) WithErrorTrace() func(*ScrollRequest) { _ = "STUB: not implemented"; return nil }

func (f Scroll) WithFilterPath(v ...string) func(*ScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Scroll) WithHeader(h map[string]string) func(*ScrollRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Scroll) WithOpaqueID(s string) func(*ScrollRequest) { _ = "STUB: not implemented"; return nil }
