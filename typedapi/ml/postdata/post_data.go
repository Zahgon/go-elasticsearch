package postdata

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
	jobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostData struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPostData func(jobid string) *PostData

func NewPostDataFunc(tp elastictransport.Interface) NewPostData {
	_ = "STUB: not implemented"
	return *new(NewPostData)
}

func New(tp elastictransport.Interface) *PostData { _ = "STUB: not implemented"; return nil }

func (r *PostData) Raw(raw io.Reader) *PostData { _ = "STUB: not implemented"; return nil }

func (r *PostData) Request(req *Request) *PostData { _ = "STUB: not implemented"; return nil }

func (r *PostData) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostData) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostData) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PostData) Header(key, value string) *PostData { _ = "STUB: not implemented"; return nil }

func (r *PostData) _jobid(jobid string) *PostData { _ = "STUB: not implemented"; return nil }

func (r *PostData) ResetEnd(datetime string) *PostData { _ = "STUB: not implemented"; return nil }

func (r *PostData) ResetStart(datetime string) *PostData { _ = "STUB: not implemented"; return nil }

func (r *PostData) ErrorTrace(errortrace bool) *PostData { _ = "STUB: not implemented"; return nil }

func (r *PostData) FilterPath(filterpaths ...string) *PostData {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostData) Human(human bool) *PostData { _ = "STUB: not implemented"; return nil }

func (r *PostData) Pretty(pretty bool) *PostData { _ = "STUB: not implemented"; return nil }
