package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSnapshotGetRepositoryFunc(t Transport) SnapshotGetRepository {
	_ = "STUB: not implemented"
	return *new(SnapshotGetRepository)
}

type SnapshotGetRepository func(o ...func(*SnapshotGetRepositoryRequest)) (*Response, error)

type SnapshotGetRepositoryRequest struct {
	Repository []string

	Local         *bool
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SnapshotGetRepositoryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotGetRepository) WithContext(v context.Context) func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGetRepository) WithRepository(v ...string) func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGetRepository) WithLocal(v bool) func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGetRepository) WithMasterTimeout(v time.Duration) func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGetRepository) WithPretty() func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGetRepository) WithHuman() func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGetRepository) WithErrorTrace() func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGetRepository) WithFilterPath(v ...string) func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGetRepository) WithHeader(h map[string]string) func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotGetRepository) WithOpaqueID(s string) func(*SnapshotGetRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}
