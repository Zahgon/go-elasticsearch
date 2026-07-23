package esapi

import (
	"context"
	"net/http"
)

func newSynonymsGetSynonymFunc(t Transport) SynonymsGetSynonym {
	_ = "STUB: not implemented"
	return *new(SynonymsGetSynonym)
}

type SynonymsGetSynonym func(id string, o ...func(*SynonymsGetSynonymRequest)) (*Response, error)

type SynonymsGetSynonymRequest struct {
	DocumentID string

	From *int
	Size *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SynonymsGetSynonymRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SynonymsGetSynonym) WithContext(v context.Context) func(*SynonymsGetSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonym) WithFrom(v int) func(*SynonymsGetSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonym) WithSize(v int) func(*SynonymsGetSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonym) WithPretty() func(*SynonymsGetSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonym) WithHuman() func(*SynonymsGetSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonym) WithErrorTrace() func(*SynonymsGetSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonym) WithFilterPath(v ...string) func(*SynonymsGetSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonym) WithHeader(h map[string]string) func(*SynonymsGetSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonym) WithOpaqueID(s string) func(*SynonymsGetSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}
