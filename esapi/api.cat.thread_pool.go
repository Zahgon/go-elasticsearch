package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCatThreadPoolFunc(t Transport) CatThreadPool {
	_ = "STUB: not implemented"
	return *new(CatThreadPool)
}

type CatThreadPool func(o ...func(*CatThreadPoolRequest)) (*Response, error)

type CatThreadPoolRequest struct {
	ThreadPoolPatterns []string

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

func (r CatThreadPoolRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatThreadPool) WithContext(v context.Context) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithThreadPoolPatterns(v ...string) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithBytes(v string) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithFormat(v string) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithH(v ...string) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithHelp(v bool) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithLocal(v bool) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithMasterTimeout(v time.Duration) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithS(v ...string) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithTime(v string) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithV(v bool) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithPretty() func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithHuman() func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithErrorTrace() func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithFilterPath(v ...string) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithHeader(h map[string]string) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatThreadPool) WithOpaqueID(s string) func(*CatThreadPoolRequest) {
	_ = "STUB: not implemented"
	return nil
}
