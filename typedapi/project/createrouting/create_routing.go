package createrouting

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
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateRouting struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCreateRouting func(name string) *CreateRouting

func NewCreateRoutingFunc(tp elastictransport.Interface) NewCreateRouting {
	_ = "STUB: not implemented"
	return *new(NewCreateRouting)
}

func New(tp elastictransport.Interface) *CreateRouting { _ = "STUB: not implemented"; return nil }

func (r *CreateRouting) Raw(raw io.Reader) *CreateRouting { _ = "STUB: not implemented"; return nil }

func (r *CreateRouting) Request(req *Request) *CreateRouting { _ = "STUB: not implemented"; return nil }

func (r *CreateRouting) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateRouting) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateRouting) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CreateRouting) Header(key, value string) *CreateRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRouting) _name(name string) *CreateRouting { _ = "STUB: not implemented"; return nil }

func (r *CreateRouting) ErrorTrace(errortrace bool) *CreateRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRouting) FilterPath(filterpaths ...string) *CreateRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRouting) Human(human bool) *CreateRouting { _ = "STUB: not implemented"; return nil }

func (r *CreateRouting) Pretty(pretty bool) *CreateRouting { _ = "STUB: not implemented"; return nil }

func (r *CreateRouting) Expression(routingexpression string) *CreateRouting {
	_ = "STUB: not implemented"
	return nil
}
