package renderquery

import (
	gobytes "bytes"
	"context"
	"encoding/json"
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

type RenderQuery struct {
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

type NewRenderQuery func(name string) *RenderQuery

func NewRenderQueryFunc(tp elastictransport.Interface) NewRenderQuery {
	_ = "STUB: not implemented"
	return *new(NewRenderQuery)
}

func New(tp elastictransport.Interface) *RenderQuery { _ = "STUB: not implemented"; return nil }

func (r *RenderQuery) Raw(raw io.Reader) *RenderQuery { _ = "STUB: not implemented"; return nil }

func (r *RenderQuery) Request(req *Request) *RenderQuery { _ = "STUB: not implemented"; return nil }

func (r *RenderQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RenderQuery) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RenderQuery) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RenderQuery) Header(key, value string) *RenderQuery { _ = "STUB: not implemented"; return nil }

func (r *RenderQuery) _name(name string) *RenderQuery { _ = "STUB: not implemented"; return nil }

func (r *RenderQuery) ErrorTrace(errortrace bool) *RenderQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderQuery) FilterPath(filterpaths ...string) *RenderQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderQuery) Human(human bool) *RenderQuery { _ = "STUB: not implemented"; return nil }

func (r *RenderQuery) Pretty(pretty bool) *RenderQuery { _ = "STUB: not implemented"; return nil }

func (r *RenderQuery) Params(params map[string]json.RawMessage) *RenderQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderQuery) AddParam(key string, value json.RawMessage) *RenderQuery {
	_ = "STUB: not implemented"
	return nil
}
