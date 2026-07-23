package postsecret

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

type PostSecret struct {
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

type NewPostSecret func() *PostSecret

func NewPostSecretFunc(tp elastictransport.Interface) NewPostSecret {
	_ = "STUB: not implemented"
	return *new(NewPostSecret)
}

func New(tp elastictransport.Interface) *PostSecret { _ = "STUB: not implemented"; return nil }

func (r *PostSecret) Raw(raw io.Reader) *PostSecret { _ = "STUB: not implemented"; return nil }

func (r *PostSecret) Request(req *Request) *PostSecret { _ = "STUB: not implemented"; return nil }

func (r *PostSecret) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostSecret) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostSecret) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PostSecret) Header(key, value string) *PostSecret { _ = "STUB: not implemented"; return nil }

func (r *PostSecret) ErrorTrace(errortrace bool) *PostSecret { _ = "STUB: not implemented"; return nil }

func (r *PostSecret) FilterPath(filterpaths ...string) *PostSecret {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostSecret) Human(human bool) *PostSecret { _ = "STUB: not implemented"; return nil }

func (r *PostSecret) Pretty(pretty bool) *PostSecret { _ = "STUB: not implemented"; return nil }

func (r *PostSecret) Value(value string) *PostSecret { _ = "STUB: not implemented"; return nil }
