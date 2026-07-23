package esapi

import (
	"context"
	"net/http"
)

func newIngestGetIPLocationDatabaseFunc(t Transport) IngestGetIPLocationDatabase {
	_ = "STUB: not implemented"
	return *new(IngestGetIPLocationDatabase)
}

type IngestGetIPLocationDatabase func(o ...func(*IngestGetIPLocationDatabaseRequest)) (*Response, error)

type IngestGetIPLocationDatabaseRequest struct {
	DocumentID []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IngestGetIPLocationDatabaseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestGetIPLocationDatabase) WithContext(v context.Context) func(*IngestGetIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetIPLocationDatabase) WithDocumentID(v ...string) func(*IngestGetIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetIPLocationDatabase) WithPretty() func(*IngestGetIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetIPLocationDatabase) WithHuman() func(*IngestGetIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetIPLocationDatabase) WithErrorTrace() func(*IngestGetIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetIPLocationDatabase) WithFilterPath(v ...string) func(*IngestGetIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetIPLocationDatabase) WithHeader(h map[string]string) func(*IngestGetIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetIPLocationDatabase) WithOpaqueID(s string) func(*IngestGetIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}
