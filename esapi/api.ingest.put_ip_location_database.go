package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIngestPutIPLocationDatabaseFunc(t Transport) IngestPutIPLocationDatabase {
	_ = "STUB: not implemented"
	return *new(IngestPutIPLocationDatabase)
}

type IngestPutIPLocationDatabase func(id string, body io.Reader, o ...func(*IngestPutIPLocationDatabaseRequest)) (*Response, error)

type IngestPutIPLocationDatabaseRequest struct {
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

func (r IngestPutIPLocationDatabaseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestPutIPLocationDatabase) WithContext(v context.Context) func(*IngestPutIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutIPLocationDatabase) WithMasterTimeout(v time.Duration) func(*IngestPutIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutIPLocationDatabase) WithTimeout(v time.Duration) func(*IngestPutIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutIPLocationDatabase) WithPretty() func(*IngestPutIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutIPLocationDatabase) WithHuman() func(*IngestPutIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutIPLocationDatabase) WithErrorTrace() func(*IngestPutIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutIPLocationDatabase) WithFilterPath(v ...string) func(*IngestPutIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutIPLocationDatabase) WithHeader(h map[string]string) func(*IngestPutIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestPutIPLocationDatabase) WithOpaqueID(s string) func(*IngestPutIPLocationDatabaseRequest) {
	_ = "STUB: not implemented"
	return nil
}
