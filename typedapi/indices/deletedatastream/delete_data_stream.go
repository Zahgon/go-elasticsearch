package deletedatastream

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

type DeleteDataStream struct {
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

type NewDeleteDataStream func(name string) *DeleteDataStream

func NewDeleteDataStreamFunc(tp elastictransport.Interface) NewDeleteDataStream {
	_ = "STUB: not implemented"
	return *new(NewDeleteDataStream)
}

func New(tp elastictransport.Interface) *DeleteDataStream { _ = "STUB: not implemented"; return nil }

func (r *DeleteDataStream) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataStream) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataStream) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataStream) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteDataStream) Header(key, value string) *DeleteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStream) _name(name string) *DeleteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStream) MasterTimeout(duration string) *DeleteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStream) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *DeleteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStream) ErrorTrace(errortrace bool) *DeleteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStream) FilterPath(filterpaths ...string) *DeleteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStream) Human(human bool) *DeleteDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataStream) Pretty(pretty bool) *DeleteDataStream {
	_ = "STUB: not implemented"
	return nil
}
