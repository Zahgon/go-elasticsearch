package search

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

type Search struct {
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

type NewSearch func(name string) *Search

func NewSearchFunc(tp elastictransport.Interface) NewSearch {
	_ = "STUB: not implemented"
	return *new(NewSearch)
}

func New(tp elastictransport.Interface) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Raw(raw io.Reader) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Request(req *Request) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Search) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Search) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Search) Header(key, value string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) _name(name string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) TypedKeys(typedkeys bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) ErrorTrace(errortrace bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) FilterPath(filterpaths ...string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Human(human bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Pretty(pretty bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Params(params map[string]json.RawMessage) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) AddParam(key string, value json.RawMessage) *Search {
	_ = "STUB: not implemented"
	return nil
}
