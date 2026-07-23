package getwatch

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

type GetWatch struct {
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

type NewGetWatch func(id string) *GetWatch

func NewGetWatchFunc(tp elastictransport.Interface) NewGetWatch {
	_ = "STUB: not implemented"
	return *new(NewGetWatch)
}

func New(tp elastictransport.Interface) *GetWatch { _ = "STUB: not implemented"; return nil }

func (r *GetWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetWatch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetWatch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetWatch) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetWatch) Header(key, value string) *GetWatch { _ = "STUB: not implemented"; return nil }

func (r *GetWatch) _id(id string) *GetWatch { _ = "STUB: not implemented"; return nil }

func (r *GetWatch) ErrorTrace(errortrace bool) *GetWatch { _ = "STUB: not implemented"; return nil }

func (r *GetWatch) FilterPath(filterpaths ...string) *GetWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetWatch) Human(human bool) *GetWatch { _ = "STUB: not implemented"; return nil }

func (r *GetWatch) Pretty(pretty bool) *GetWatch { _ = "STUB: not implemented"; return nil }
