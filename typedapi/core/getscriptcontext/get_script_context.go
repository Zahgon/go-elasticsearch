package getscriptcontext

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetScriptContext struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetScriptContext func() *GetScriptContext

func NewGetScriptContextFunc(tp elastictransport.Interface) NewGetScriptContext {
	_ = "STUB: not implemented"
	return *new(NewGetScriptContext)
}

func New(tp elastictransport.Interface) *GetScriptContext { _ = "STUB: not implemented"; return nil }

func (r *GetScriptContext) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetScriptContext) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetScriptContext) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetScriptContext) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetScriptContext) Header(key, value string) *GetScriptContext {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScriptContext) ErrorTrace(errortrace bool) *GetScriptContext {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScriptContext) FilterPath(filterpaths ...string) *GetScriptContext {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScriptContext) Human(human bool) *GetScriptContext {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScriptContext) Pretty(pretty bool) *GetScriptContext {
	_ = "STUB: not implemented"
	return nil
}
