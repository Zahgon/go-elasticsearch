package getrouting

import (
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

type GetRouting struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetRouting func(name string) *GetRouting

func NewGetRoutingFunc(tp elastictransport.Interface) NewGetRouting {
	_ = "STUB: not implemented"
	return *new(NewGetRouting)
}

func New(tp elastictransport.Interface) *GetRouting { _ = "STUB: not implemented"; return nil }

func (r *GetRouting) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRouting) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRouting) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRouting) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetRouting) Header(key, value string) *GetRouting { _ = "STUB: not implemented"; return nil }

func (r *GetRouting) _name(name string) *GetRouting { _ = "STUB: not implemented"; return nil }

func (r *GetRouting) ErrorTrace(errortrace bool) *GetRouting { _ = "STUB: not implemented"; return nil }

func (r *GetRouting) FilterPath(filterpaths ...string) *GetRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRouting) Human(human bool) *GetRouting { _ = "STUB: not implemented"; return nil }

func (r *GetRouting) Pretty(pretty bool) *GetRouting { _ = "STUB: not implemented"; return nil }
