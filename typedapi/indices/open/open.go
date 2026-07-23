package open

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
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Open struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewOpen func(index string) *Open

func NewOpenFunc(tp elastictransport.Interface) NewOpen {
	_ = "STUB: not implemented"
	return *new(NewOpen)
}

func New(tp elastictransport.Interface) *Open { _ = "STUB: not implemented"; return nil }

func (r *Open) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Open) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Open) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Open) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Open) Header(key, value string) *Open { _ = "STUB: not implemented"; return nil }

func (r *Open) _index(index string) *Open { _ = "STUB: not implemented"; return nil }

func (r *Open) AllowNoIndices(allownoindices bool) *Open { _ = "STUB: not implemented"; return nil }

func (r *Open) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Open {
	_ = "STUB: not implemented"
	return nil
}

func (r *Open) IgnoreUnavailable(ignoreunavailable bool) *Open {
	_ = "STUB: not implemented"
	return nil
}

func (r *Open) MasterTimeout(duration string) *Open { _ = "STUB: not implemented"; return nil }

func (r *Open) Timeout(duration string) *Open { _ = "STUB: not implemented"; return nil }

func (r *Open) WaitForActiveShards(waitforactiveshards string) *Open {
	_ = "STUB: not implemented"
	return nil
}

func (r *Open) ErrorTrace(errortrace bool) *Open { _ = "STUB: not implemented"; return nil }

func (r *Open) FilterPath(filterpaths ...string) *Open { _ = "STUB: not implemented"; return nil }

func (r *Open) Human(human bool) *Open { _ = "STUB: not implemented"; return nil }

func (r *Open) Pretty(pretty bool) *Open { _ = "STUB: not implemented"; return nil }
