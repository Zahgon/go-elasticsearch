package getdatastreammappings

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

type GetDataStreamMappings struct {
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

type NewGetDataStreamMappings func(name string) *GetDataStreamMappings

func NewGetDataStreamMappingsFunc(tp elastictransport.Interface) NewGetDataStreamMappings {
	_ = "STUB: not implemented"
	return *new(NewGetDataStreamMappings)
}

func New(tp elastictransport.Interface) *GetDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamMappings) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStreamMappings) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStreamMappings) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStreamMappings) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDataStreamMappings) Header(key, value string) *GetDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamMappings) _name(name string) *GetDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamMappings) MasterTimeout(duration string) *GetDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamMappings) ErrorTrace(errortrace bool) *GetDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamMappings) FilterPath(filterpaths ...string) *GetDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamMappings) Human(human bool) *GetDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamMappings) Pretty(pretty bool) *GetDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}
