package esapi

import (
	"context"
	"net/http"
)

func newIngestGetGeoipDatabaseFunc(t Transport) IngestGetGeoipDatabase {
	_ = "STUB: not implemented"
	return *new(IngestGetGeoipDatabase)
}

type IngestGetGeoipDatabase func(o ...func(*IngestGetGeoipDatabaseRequest)) (*Response, error)

type IngestGetGeoipDatabaseRequest struct {
	DocumentID []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IngestGetGeoipDatabaseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestGetGeoipDatabase) WithContext(v context.Context) func(*IngestGetGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetGeoipDatabase) WithDocumentID(v ...string) func(*IngestGetGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetGeoipDatabase) WithPretty() func(*IngestGetGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetGeoipDatabase) WithHuman() func(*IngestGetGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetGeoipDatabase) WithErrorTrace() func(*IngestGetGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetGeoipDatabase) WithFilterPath(v ...string) func(*IngestGetGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetGeoipDatabase) WithHeader(h map[string]string) func(*IngestGetGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGetGeoipDatabase) WithOpaqueID(s string) func(*IngestGetGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}
