package getsecret

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

type GetSecret struct {
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

type NewGetSecret func(id string) *GetSecret

func NewGetSecretFunc(tp elastictransport.Interface) NewGetSecret {
	_ = "STUB: not implemented"
	return *new(NewGetSecret)
}

func New(tp elastictransport.Interface) *GetSecret { _ = "STUB: not implemented"; return nil }

func (r *GetSecret) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSecret) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSecret) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSecret) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetSecret) Header(key, value string) *GetSecret { _ = "STUB: not implemented"; return nil }

func (r *GetSecret) _id(id string) *GetSecret { _ = "STUB: not implemented"; return nil }

func (r *GetSecret) ErrorTrace(errortrace bool) *GetSecret { _ = "STUB: not implemented"; return nil }

func (r *GetSecret) FilterPath(filterpaths ...string) *GetSecret {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSecret) Human(human bool) *GetSecret { _ = "STUB: not implemented"; return nil }

func (r *GetSecret) Pretty(pretty bool) *GetSecret { _ = "STUB: not implemented"; return nil }
