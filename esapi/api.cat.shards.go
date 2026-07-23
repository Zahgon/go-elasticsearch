package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatShardsFunc(t Transport) CatShards { _ = "STUB: not implemented"; return *new(CatShards) }

type CatShards func(o ...func(*CatShardsRequest)) (*Response, error)

type CatShardsRequest struct {
	Index []string

	Bytes         string
	Format        string
	H             []string
	Help          *bool
	MasterTimeout time.Duration
	S             []string
	Time          string
	V             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatShardsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatShards) WithContext(v context.Context) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithIndex(v ...string) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithBytes(v string) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithFormat(v string) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithH(v ...string) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithHelp(v bool) func(*CatShardsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatShards) WithMasterTimeout(v time.Duration) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithS(v ...string) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithTime(v string) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithV(v bool) func(*CatShardsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatShards) WithPretty() func(*CatShardsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatShards) WithHuman() func(*CatShardsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatShards) WithErrorTrace() func(*CatShardsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatShards) WithFilterPath(v ...string) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithHeader(h map[string]string) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatShards) WithOpaqueID(s string) func(*CatShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}
