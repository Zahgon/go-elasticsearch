package esapi

import (
	"context"
	"net/http"
)

func newSynonymsDeleteSynonymFunc(t Transport) SynonymsDeleteSynonym {
	_ = "STUB: not implemented"
	return *new(SynonymsDeleteSynonym)
}

type SynonymsDeleteSynonym func(id string, o ...func(*SynonymsDeleteSynonymRequest)) (*Response, error)

type SynonymsDeleteSynonymRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SynonymsDeleteSynonymRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SynonymsDeleteSynonym) WithContext(v context.Context) func(*SynonymsDeleteSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonym) WithPretty() func(*SynonymsDeleteSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonym) WithHuman() func(*SynonymsDeleteSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonym) WithErrorTrace() func(*SynonymsDeleteSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonym) WithFilterPath(v ...string) func(*SynonymsDeleteSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonym) WithHeader(h map[string]string) func(*SynonymsDeleteSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonym) WithOpaqueID(s string) func(*SynonymsDeleteSynonymRequest) {
	_ = "STUB: not implemented"
	return nil
}
