package deletedatastreamoptions

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

type DeleteDataStreamOptions struct {
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

type NewDeleteDataStreamOptions func(name string) *DeleteDataStreamOptions

func NewDeleteDataStreamOptionsFunc(tp elastictransport.Interface) NewDeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return *new(NewDeleteDataStreamOptions)
}

func New(tp elastictransport.Interface) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStreamOptions) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataStreamOptions) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataStreamOptions) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataStreamOptions) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteDataStreamOptions) Header(key, value string) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStreamOptions) _name(name string) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStreamOptions) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStreamOptions) MasterTimeout(duration string) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStreamOptions) Timeout(duration string) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStreamOptions) ErrorTrace(errortrace bool) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStreamOptions) FilterPath(filterpaths ...string) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStreamOptions) Human(human bool) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStreamOptions) Pretty(pretty bool) *DeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return nil
}
