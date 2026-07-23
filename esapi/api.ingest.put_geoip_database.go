package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIngestPutGeoipDatabaseFunc(t Transport) IngestPutGeoipDatabase {
	_ = "STUB: not implemented"
	return *new(IngestPutGeoipDatabase)
}

type IngestPutGeoipDatabase func(id string, body io.Reader, o ...func(*IngestPutGeoipDatabaseRequest)) (*Response, error)

type IngestPutGeoipDatabaseRequest struct {
	DocumentID string

	Body io.Reader

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

func (r IngestPutGeoipDatabaseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestPutGeoipDatabase) WithContext(v context.Context) func(*IngestPutGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutGeoipDatabase) WithMasterTimeout(v time.Duration) func(*IngestPutGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutGeoipDatabase) WithTimeout(v time.Duration) func(*IngestPutGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutGeoipDatabase) WithPretty() func(*IngestPutGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutGeoipDatabase) WithHuman() func(*IngestPutGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutGeoipDatabase) WithErrorTrace() func(*IngestPutGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutGeoipDatabase) WithFilterPath(v ...string) func(*IngestPutGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutGeoipDatabase) WithHeader(h map[string]string) func(*IngestPutGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutGeoipDatabase) WithOpaqueID(s string) func(*IngestPutGeoipDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}
