package esapi

import (
	"context"
	"net/http"
)

func newSnapshotRepositoryVerifyIntegrityFunc(t Transport) SnapshotRepositoryVerifyIntegrity {
	_ = "STUB: not implemented"
	return *new(SnapshotRepositoryVerifyIntegrity)
}

type SnapshotRepositoryVerifyIntegrity func(repository []string, o ...func(*SnapshotRepositoryVerifyIntegrityRequest)) (*Response, error)

type SnapshotRepositoryVerifyIntegrityRequest struct {
	Repository []string

	BlobThreadPoolConcurrency            *int
	IndexSnapshotVerificationConcurrency *int
	IndexVerificationConcurrency         *int
	MaxBytesPerSec                       string
	MaxFailedShardSnapshots              *int
	MetaThreadPoolConcurrency            *int
	SnapshotVerificationConcurrency      *int
	VerifyBlobContents                   *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SnapshotRepositoryVerifyIntegrityRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithContext(v context.Context) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithBlobThreadPoolConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithIndexSnapshotVerificationConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithIndexVerificationConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithMaxBytesPerSec(v string) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithMaxFailedShardSnapshots(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithMetaThreadPoolConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithSnapshotVerificationConcurrency(v int) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithVerifyBlobContents(v bool) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithPretty() func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithHuman() func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithErrorTrace() func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithFilterPath(v ...string) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithHeader(h map[string]string) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SnapshotRepositoryVerifyIntegrity) WithOpaqueID(s string) func(*SnapshotRepositoryVerifyIntegrityRequest) {
	_ = "STUB: not implemented"
	return nil
}
