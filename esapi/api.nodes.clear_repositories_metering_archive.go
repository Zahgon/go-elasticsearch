package esapi

import (
	"context"
	"net/http"
)

func newNodesClearRepositoriesMeteringArchiveFunc(t Transport) NodesClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return *new(NodesClearRepositoriesMeteringArchive)
}

type NodesClearRepositoriesMeteringArchive func(max_archive_version *int64, node_id []string, o ...func(*NodesClearRepositoriesMeteringArchiveRequest)) (*Response, error)

type NodesClearRepositoriesMeteringArchiveRequest struct {
	MaxArchiveVersion *int64
	NodeID            []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r NodesClearRepositoriesMeteringArchiveRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f NodesClearRepositoriesMeteringArchive) WithContext(v context.Context) func(*NodesClearRepositoriesMeteringArchiveRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesClearRepositoriesMeteringArchive) WithPretty() func(*NodesClearRepositoriesMeteringArchiveRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesClearRepositoriesMeteringArchive) WithHuman() func(*NodesClearRepositoriesMeteringArchiveRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesClearRepositoriesMeteringArchive) WithErrorTrace() func(*NodesClearRepositoriesMeteringArchiveRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesClearRepositoriesMeteringArchive) WithFilterPath(v ...string) func(*NodesClearRepositoriesMeteringArchiveRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesClearRepositoriesMeteringArchive) WithHeader(h map[string]string) func(*NodesClearRepositoriesMeteringArchiveRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesClearRepositoriesMeteringArchive) WithOpaqueID(s string) func(*NodesClearRepositoriesMeteringArchiveRequest) {
	_ = "STUB: not implemented"
	return nil
}
