package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSnapshotCleanupRepositoryFunc(t Transport) SnapshotCleanupRepository {
	_ = "STUB: not implemented"
	return *new(SnapshotCleanupRepository)
}

type SnapshotCleanupRepository func(repository string, o ...func(*SnapshotCleanupRepositoryRequest)) (*Response, error)

type SnapshotCleanupRepositoryRequest struct {
	Repository string

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

func (r SnapshotCleanupRepositoryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotCleanupRepository) WithContext(v context.Context) func(*SnapshotCleanupRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCleanupRepository) WithMasterTimeout(v time.Duration) func(*SnapshotCleanupRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCleanupRepository) WithTimeout(v time.Duration) func(*SnapshotCleanupRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCleanupRepository) WithPretty() func(*SnapshotCleanupRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCleanupRepository) WithHuman() func(*SnapshotCleanupRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCleanupRepository) WithErrorTrace() func(*SnapshotCleanupRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCleanupRepository) WithFilterPath(v ...string) func(*SnapshotCleanupRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCleanupRepository) WithHeader(h map[string]string) func(*SnapshotCleanupRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotCleanupRepository) WithOpaqueID(s string) func(*SnapshotCleanupRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}
