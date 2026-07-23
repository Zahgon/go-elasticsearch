package createmanyrouting

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

type CreateManyRouting struct {
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

type NewCreateManyRouting func() *CreateManyRouting

func NewCreateManyRoutingFunc(tp elastictransport.Interface) NewCreateManyRouting {
	_ = "STUB: not implemented"
	return *new(NewCreateManyRouting)
}

func New(tp elastictransport.Interface) *CreateManyRouting { _ = "STUB: not implemented"; return nil }

func (r *CreateManyRouting) Raw(raw io.Reader) *CreateManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateManyRouting) Request(req *Request) *CreateManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateManyRouting) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateManyRouting) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateManyRouting) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CreateManyRouting) Header(key, value string) *CreateManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateManyRouting) ErrorTrace(errortrace bool) *CreateManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateManyRouting) FilterPath(filterpaths ...string) *CreateManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateManyRouting) Human(human bool) *CreateManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateManyRouting) Pretty(pretty bool) *CreateManyRouting {
	_ = "STUB: not implemented"
	return nil
}
