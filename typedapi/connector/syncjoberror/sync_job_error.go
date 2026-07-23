package syncjoberror

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	connectorsyncjobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SyncJobError struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorsyncjobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSyncJobError func(connectorsyncjobid string) *SyncJobError

func NewSyncJobErrorFunc(tp elastictransport.Interface) NewSyncJobError {
	_ = "STUB: not implemented"
	return *new(NewSyncJobError)
}

func New(tp elastictransport.Interface) *SyncJobError { _ = "STUB: not implemented"; return nil }

func (r *SyncJobError) Raw(raw io.Reader) *SyncJobError { _ = "STUB: not implemented"; return nil }

func (r *SyncJobError) Request(req *Request) *SyncJobError { _ = "STUB: not implemented"; return nil }

func (r *SyncJobError) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobError) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobError) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SyncJobError) Header(key, value string) *SyncJobError {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobError) _connectorsyncjobid(connectorsyncjobid string) *SyncJobError {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobError) ErrorTrace(errortrace bool) *SyncJobError {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobError) FilterPath(filterpaths ...string) *SyncJobError {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobError) Human(human bool) *SyncJobError { _ = "STUB: not implemented"; return nil }

func (r *SyncJobError) Pretty(pretty bool) *SyncJobError { _ = "STUB: not implemented"; return nil }

func (r *SyncJobError) Error(error string) *SyncJobError { _ = "STUB: not implemented"; return nil }
