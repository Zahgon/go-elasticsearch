package clone

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
	repositoryMask = iota + 1

	snapshotMask

	targetsnapshotMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Clone struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	repository     string
	snapshot       string
	targetsnapshot string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewClone func(repository, snapshot, targetsnapshot string) *Clone

func NewCloneFunc(tp elastictransport.Interface) NewClone {
	_ = "STUB: not implemented"
	return *new(NewClone)
}

func New(tp elastictransport.Interface) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) Raw(raw io.Reader) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) Request(req *Request) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Clone) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Clone) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Clone) Header(key, value string) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) _repository(repository string) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) _snapshot(snapshot string) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) _targetsnapshot(targetsnapshot string) *Clone {
	_ = "STUB: not implemented"
	return nil
}

func (r *Clone) MasterTimeout(duration string) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) ErrorTrace(errortrace bool) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) FilterPath(filterpaths ...string) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) Human(human bool) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) Pretty(pretty bool) *Clone { _ = "STUB: not implemented"; return nil }

func (r *Clone) Indices(indices string) *Clone { _ = "STUB: not implemented"; return nil }
