package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIngestDeleteIPLocationDatabaseFunc(t Transport) IngestDeleteIPLocationDatabase {
	_ = "STUB: not implemented"
	return *new(IngestDeleteIPLocationDatabase)
}

type IngestDeleteIPLocationDatabase func(id []string, o ...func(*IngestDeleteIPLocationDatabaseRequest)) (*Response, error)

type IngestDeleteIPLocationDatabaseRequest struct {
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

func (r IngestDeleteIPLocationDatabaseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestDeleteIPLocationDatabase) WithContext(v context.Context) func(*IngestDeleteIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteIPLocationDatabase) WithMasterTimeout(v time.Duration) func(*IngestDeleteIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteIPLocationDatabase) WithTimeout(v time.Duration) func(*IngestDeleteIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteIPLocationDatabase) WithPretty() func(*IngestDeleteIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteIPLocationDatabase) WithHuman() func(*IngestDeleteIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteIPLocationDatabase) WithErrorTrace() func(*IngestDeleteIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteIPLocationDatabase) WithFilterPath(v ...string) func(*IngestDeleteIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteIPLocationDatabase) WithHeader(h map[string]string) func(*IngestDeleteIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestDeleteIPLocationDatabase) WithOpaqueID(s string) func(*IngestDeleteIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}
