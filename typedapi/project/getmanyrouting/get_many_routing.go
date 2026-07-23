package getmanyrouting

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetManyRouting struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetManyRouting func() *GetManyRouting

func NewGetManyRoutingFunc(tp elastictransport.Interface) NewGetManyRouting {
	_ = "STUB: not implemented"
	return *new(NewGetManyRouting)
}

func New(tp elastictransport.Interface) *GetManyRouting { _ = "STUB: not implemented"; return nil }

func (r *GetManyRouting) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetManyRouting) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetManyRouting) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetManyRouting) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetManyRouting) Header(key, value string) *GetManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetManyRouting) ErrorTrace(errortrace bool) *GetManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetManyRouting) FilterPath(filterpaths ...string) *GetManyRouting {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetManyRouting) Human(human bool) *GetManyRouting { _ = "STUB: not implemented"; return nil }

func (r *GetManyRouting) Pretty(pretty bool) *GetManyRouting { _ = "STUB: not implemented"; return nil }
