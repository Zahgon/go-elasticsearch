package getdatastreamoptions

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetDataStreamOptions struct {
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

type NewGetDataStreamOptions func(name string) *GetDataStreamOptions

func NewGetDataStreamOptionsFunc(tp elastictransport.Interface) NewGetDataStreamOptions {
	_ = "STUB: not implemented"
	return *new(NewGetDataStreamOptions)
}

func New(tp elastictransport.Interface) *GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamOptions) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStreamOptions) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStreamOptions) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStreamOptions) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDataStreamOptions) Header(key, value string) *GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamOptions) _name(name string) *GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamOptions) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamOptions) MasterTimeout(duration string) *GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamOptions) ErrorTrace(errortrace bool) *GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamOptions) FilterPath(filterpaths ...string) *GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamOptions) Human(human bool) *GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStreamOptions) Pretty(pretty bool) *GetDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}
