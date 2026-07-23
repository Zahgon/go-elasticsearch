package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSnapshotDeleteRepositoryFunc(t Transport) SnapshotDeleteRepository {
	_ = "STUB: not implemented"
	return *new(SnapshotDeleteRepository)
}

type SnapshotDeleteRepository func(repository []string, o ...func(*SnapshotDeleteRepositoryRequest)) (*Response, error)

type SnapshotDeleteRepositoryRequest struct {
	Repository []string

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

func (r SnapshotDeleteRepositoryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotDeleteRepository) WithContext(v context.Context) func(*SnapshotDeleteRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDeleteRepository) WithMasterTimeout(v time.Duration) func(*SnapshotDeleteRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDeleteRepository) WithTimeout(v time.Duration) func(*SnapshotDeleteRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDeleteRepository) WithPretty() func(*SnapshotDeleteRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDeleteRepository) WithHuman() func(*SnapshotDeleteRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDeleteRepository) WithErrorTrace() func(*SnapshotDeleteRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDeleteRepository) WithFilterPath(v ...string) func(*SnapshotDeleteRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDeleteRepository) WithHeader(h map[string]string) func(*SnapshotDeleteRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotDeleteRepository) WithOpaqueID(s string) func(*SnapshotDeleteRepositoryRequest) {
	_ = "STUB: not implemented"
	return nil
}
