package getdatastream

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

type GetDataStream struct {
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

type NewGetDataStream func() *GetDataStream

func NewGetDataStreamFunc(tp elastictransport.Interface) NewGetDataStream {
	_ = "STUB: not implemented"
	return *new(NewGetDataStream)
}

func New(tp elastictransport.Interface) *GetDataStream { _ = "STUB: not implemented"; return nil }

func (r *GetDataStream) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStream) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStream) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataStream) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDataStream) Header(key, value string) *GetDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStream) Name(name string) *GetDataStream { _ = "STUB: not implemented"; return nil }

func (r *GetDataStream) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStream) IncludeDefaults(includedefaults bool) *GetDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStream) MasterTimeout(duration string) *GetDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStream) Verbose(verbose bool) *GetDataStream { _ = "STUB: not implemented"; return nil }

func (r *GetDataStream) ErrorTrace(errortrace bool) *GetDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStream) FilterPath(filterpaths ...string) *GetDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataStream) Human(human bool) *GetDataStream { _ = "STUB: not implemented"; return nil }

func (r *GetDataStream) Pretty(pretty bool) *GetDataStream { _ = "STUB: not implemented"; return nil }
