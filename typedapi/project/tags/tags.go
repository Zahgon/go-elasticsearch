package tags

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Tags struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewTags func() *Tags

func NewTagsFunc(tp elastictransport.Interface) NewTags {
	_ = "STUB: not implemented"
	return *new(NewTags)
}

func New(tp elastictransport.Interface) *Tags { _ = "STUB: not implemented"; return nil }

func (r *Tags) Raw(raw io.Reader) *Tags { _ = "STUB: not implemented"; return nil }

func (r *Tags) Request(req *Request) *Tags { _ = "STUB: not implemented"; return nil }

func (r *Tags) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Tags) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Tags) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Tags) Header(key, value string) *Tags { _ = "STUB: not implemented"; return nil }

func (r *Tags) ErrorTrace(errortrace bool) *Tags { _ = "STUB: not implemented"; return nil }

func (r *Tags) FilterPath(filterpaths ...string) *Tags { _ = "STUB: not implemented"; return nil }

func (r *Tags) Human(human bool) *Tags { _ = "STUB: not implemented"; return nil }

func (r *Tags) Pretty(pretty bool) *Tags { _ = "STUB: not implemented"; return nil }

func (r *Tags) ProjectRouting(projectrouting string) *Tags { _ = "STUB: not implemented"; return nil }
