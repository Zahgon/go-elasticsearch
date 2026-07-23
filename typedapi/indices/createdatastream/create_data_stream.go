package createdatastream

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

type CreateDataStream struct {
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

type NewCreateDataStream func(name string) *CreateDataStream

func NewCreateDataStreamFunc(tp elastictransport.Interface) NewCreateDataStream {
	_ = "STUB: not implemented"
	return *new(NewCreateDataStream)
}

func New(tp elastictransport.Interface) *CreateDataStream { _ = "STUB: not implemented"; return nil }

func (r *CreateDataStream) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateDataStream) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateDataStream) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateDataStream) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *CreateDataStream) Header(key, value string) *CreateDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateDataStream) _name(name string) *CreateDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateDataStream) MasterTimeout(duration string) *CreateDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateDataStream) Timeout(duration string) *CreateDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateDataStream) ErrorTrace(errortrace bool) *CreateDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateDataStream) FilterPath(filterpaths ...string) *CreateDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateDataStream) Human(human bool) *CreateDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateDataStream) Pretty(pretty bool) *CreateDataStream {
	_ = "STUB: not implemented"
	return nil
}
