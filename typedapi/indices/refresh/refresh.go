package refresh

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

type Refresh struct {
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

type NewRefresh func() *Refresh

func NewRefreshFunc(tp elastictransport.Interface) NewRefresh {
	_ = "STUB: not implemented"
	return *new(NewRefresh)
}

func New(tp elastictransport.Interface) *Refresh { _ = "STUB: not implemented"; return nil }

func (r *Refresh) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Refresh) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Refresh) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Refresh) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Refresh) Header(key, value string) *Refresh { _ = "STUB: not implemented"; return nil }

func (r *Refresh) Index(index string) *Refresh { _ = "STUB: not implemented"; return nil }

func (r *Refresh) AllowNoIndices(allownoindices bool) *Refresh {
	_ = "STUB: not implemented"
	return nil
}

func (r *Refresh) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Refresh {
	_ = "STUB: not implemented"
	return nil
}

func (r *Refresh) IgnoreUnavailable(ignoreunavailable bool) *Refresh {
	_ = "STUB: not implemented"
	return nil
}

func (r *Refresh) ErrorTrace(errortrace bool) *Refresh { _ = "STUB: not implemented"; return nil }

func (r *Refresh) FilterPath(filterpaths ...string) *Refresh { _ = "STUB: not implemented"; return nil }

func (r *Refresh) Human(human bool) *Refresh { _ = "STUB: not implemented"; return nil }

func (r *Refresh) Pretty(pretty bool) *Refresh { _ = "STUB: not implemented"; return nil }
