package esapi

import (
	"context"
	"io"
	"net/http"
)

func newWatcherPutWatchFunc(t Transport) WatcherPutWatch {
	_ = "STUB: not implemented"
	return *new(WatcherPutWatch)
}

type WatcherPutWatch func(id string, body io.Reader, o ...func(*WatcherPutWatchRequest)) (*Response, error)

type WatcherPutWatchRequest struct {
	WatchID string

	Body io.Reader

	Active        *bool
	IfPrimaryTerm *int64
	IfSeqNo       *int64
	Version       *int64

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherPutWatchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherPutWatch) WithContext(v context.Context) func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithActive(v bool) func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithIfPrimaryTerm(v int64) func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithIfSeqNo(v int64) func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithVersion(v int64) func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithPretty() func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithHuman() func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithErrorTrace() func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithFilterPath(v ...string) func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithHeader(h map[string]string) func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherPutWatch) WithOpaqueID(s string) func(*WatcherPutWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}
