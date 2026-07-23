package esapi

import (
	"context"
	"net/http"
)

func newSynonymsGetSynonymsSetsFunc(t Transport) SynonymsGetSynonymsSets {
	_ = "STUB: not implemented"
	return *new(SynonymsGetSynonymsSets)
}

type SynonymsGetSynonymsSets func(o ...func(*SynonymsGetSynonymsSetsRequest)) (*Response, error)

type SynonymsGetSynonymsSetsRequest struct {
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

func (r SynonymsGetSynonymsSetsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SynonymsGetSynonymsSets) WithContext(v context.Context) func(*SynonymsGetSynonymsSetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymsSets) WithFrom(v int) func(*SynonymsGetSynonymsSetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymsSets) WithSize(v int) func(*SynonymsGetSynonymsSetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymsSets) WithPretty() func(*SynonymsGetSynonymsSetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymsSets) WithHuman() func(*SynonymsGetSynonymsSetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymsSets) WithErrorTrace() func(*SynonymsGetSynonymsSetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymsSets) WithFilterPath(v ...string) func(*SynonymsGetSynonymsSetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymsSets) WithHeader(h map[string]string) func(*SynonymsGetSynonymsSetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymsSets) WithOpaqueID(s string) func(*SynonymsGetSynonymsSetsRequest) {
	_ = "STUB: not implemented"
	return nil
}
