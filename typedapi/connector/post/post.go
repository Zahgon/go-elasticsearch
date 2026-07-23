package post

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

type Post struct {
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

type NewPost func() *Post

func NewPostFunc(tp elastictransport.Interface) NewPost {
	_ = "STUB: not implemented"
	return *new(NewPost)
}

func New(tp elastictransport.Interface) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) Raw(raw io.Reader) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) Request(req *Request) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Post) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Post) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Post) Header(key, value string) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) ErrorTrace(errortrace bool) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) FilterPath(filterpaths ...string) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) Human(human bool) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) Pretty(pretty bool) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) Description(description string) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) IndexName(indexname string) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) IsNative(isnative bool) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) Language(language string) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) Name(name string) *Post { _ = "STUB: not implemented"; return nil }

func (r *Post) ServiceType(servicetype string) *Post { _ = "STUB: not implemented"; return nil }
