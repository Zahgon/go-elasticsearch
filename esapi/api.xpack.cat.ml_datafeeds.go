package esapi

import (
	"context"
	"net/http"
)

func newCatMLDatafeedsFunc(t Transport) CatMLDatafeeds {
	_ = "STUB: not implemented"
	return *new(CatMLDatafeeds)
}

type CatMLDatafeeds func(o ...func(*CatMLDatafeedsRequest)) (*Response, error)

type CatMLDatafeedsRequest struct {
	DatafeedID string

	AllowNoMatch *bool
	Bytes        string
	Format       string
	H            []string
	Help         *bool
	S            []string
	Time         string
	V            *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatMLDatafeedsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatMLDatafeeds) WithContext(v context.Context) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithDatafeedID(v string) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithAllowNoMatch(v bool) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithBytes(v string) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithFormat(v string) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithH(v ...string) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithHelp(v bool) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithS(v ...string) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithTime(v string) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithV(v bool) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithPretty() func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithHuman() func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithErrorTrace() func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithFilterPath(v ...string) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithHeader(h map[string]string) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDatafeeds) WithOpaqueID(s string) func(*CatMLDatafeedsRequest) {
	_ = "STUB: not implemented"
	return nil
}
