package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatRepositoriesFunc(t Transport) CatRepositories {
	_ = "STUB: not implemented"
	return *new(CatRepositories)
}

type CatRepositories func(o ...func(*CatRepositoriesRequest)) (*Response, error)

type CatRepositoriesRequest struct {
	Bytes         string
	Format        string
	H             []string
	Help          *bool
	Local         *bool
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

func (r CatRepositoriesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatRepositories) WithContext(v context.Context) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithBytes(v string) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithFormat(v string) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithH(v ...string) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithHelp(v bool) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithLocal(v bool) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithMasterTimeout(v time.Duration) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithS(v ...string) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithTime(v string) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithV(v bool) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithPretty() func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithHuman() func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithErrorTrace() func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithFilterPath(v ...string) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithHeader(h map[string]string) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatRepositories) WithOpaqueID(s string) func(*CatRepositoriesRequest) {
	_ = "STUB: not implemented"
	return nil
}
