package promotedatastream

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

type PromoteDataStream struct {
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

type NewPromoteDataStream func(name string) *PromoteDataStream

func NewPromoteDataStreamFunc(tp elastictransport.Interface) NewPromoteDataStream {
	_ = "STUB: not implemented"
	return *new(NewPromoteDataStream)
}

func New(tp elastictransport.Interface) *PromoteDataStream { _ = "STUB: not implemented"; return nil }

func (r *PromoteDataStream) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PromoteDataStream) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PromoteDataStream) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r PromoteDataStream) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PromoteDataStream) Header(key, value string) *PromoteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *PromoteDataStream) _name(name string) *PromoteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *PromoteDataStream) MasterTimeout(duration string) *PromoteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *PromoteDataStream) ErrorTrace(errortrace bool) *PromoteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *PromoteDataStream) FilterPath(filterpaths ...string) *PromoteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *PromoteDataStream) Human(human bool) *PromoteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *PromoteDataStream) Pretty(pretty bool) *PromoteDataStream {
	_ = "STUB: not implemented"
	return nil
}
