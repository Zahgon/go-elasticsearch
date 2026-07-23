package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIngestDeleteGeoipDatabaseFunc(t Transport) IngestDeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return *new(IngestDeleteGeoipDatabase)
}

type IngestDeleteGeoipDatabase func(id []string, o ...func(*IngestDeleteGeoipDatabaseRequest)) (*Response, error)

type IngestDeleteGeoipDatabaseRequest struct {
	DocumentID []string

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

func (r IngestDeleteGeoipDatabaseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestDeleteGeoipDatabase) WithContext(v context.Context) func(*IngestDeleteGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteGeoipDatabase) WithMasterTimeout(v time.Duration) func(*IngestDeleteGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteGeoipDatabase) WithTimeout(v time.Duration) func(*IngestDeleteGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteGeoipDatabase) WithPretty() func(*IngestDeleteGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteGeoipDatabase) WithHuman() func(*IngestDeleteGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteGeoipDatabase) WithErrorTrace() func(*IngestDeleteGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteGeoipDatabase) WithFilterPath(v ...string) func(*IngestDeleteGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteGeoipDatabase) WithHeader(h map[string]string) func(*IngestDeleteGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteGeoipDatabase) WithOpaqueID(s string) func(*IngestDeleteGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}
