package getscript

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetScript struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetScript func(id string) *GetScript

func NewGetScriptFunc(tp elastictransport.Interface) NewGetScript {
	_ = "STUB: not implemented"
	return *new(NewGetScript)
}

func New(tp elastictransport.Interface) *GetScript { _ = "STUB: not implemented"; return nil }

func (r *GetScript) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetScript) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetScript) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetScript) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetScript) Header(key, value string) *GetScript { _ = "STUB: not implemented"; return nil }

func (r *GetScript) _id(id string) *GetScript { _ = "STUB: not implemented"; return nil }

func (r *GetScript) MasterTimeout(duration string) *GetScript {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScript) ErrorTrace(errortrace bool) *GetScript { _ = "STUB: not implemented"; return nil }

func (r *GetScript) FilterPath(filterpaths ...string) *GetScript {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScript) Human(human bool) *GetScript { _ = "STUB: not implemented"; return nil }

func (r *GetScript) Pretty(pretty bool) *GetScript { _ = "STUB: not implemented"; return nil }
