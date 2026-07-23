package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSynonymsPutSynonymFunc(t Transport) SynonymsPutSynonym {
	_ = "STUB: not implemented"
	return *new(SynonymsPutSynonym)
}

type SynonymsPutSynonym func(id string, body io.Reader, o ...func(*SynonymsPutSynonymRequest)) (*Response, error)

type SynonymsPutSynonymRequest struct {
	DocumentID string

	Body io.Reader

	Refresh *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SynonymsPutSynonymRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SynonymsPutSynonym) WithContext(v context.Context) func(*SynonymsPutSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonym) WithRefresh(v bool) func(*SynonymsPutSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonym) WithPretty() func(*SynonymsPutSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonym) WithHuman() func(*SynonymsPutSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonym) WithErrorTrace() func(*SynonymsPutSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonym) WithFilterPath(v ...string) func(*SynonymsPutSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonym) WithHeader(h map[string]string) func(*SynonymsPutSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonym) WithOpaqueID(s string) func(*SynonymsPutSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}
