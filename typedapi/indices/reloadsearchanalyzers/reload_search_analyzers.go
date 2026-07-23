package reloadsearchanalyzers

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

type ReloadSearchAnalyzers struct {
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

type NewReloadSearchAnalyzers func(index string) *ReloadSearchAnalyzers

func NewReloadSearchAnalyzersFunc(tp elastictransport.Interface) NewReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return *new(NewReloadSearchAnalyzers)
}

func New(tp elastictransport.Interface) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReloadSearchAnalyzers) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReloadSearchAnalyzers) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ReloadSearchAnalyzers) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ReloadSearchAnalyzers) Header(key, value string) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) _index(index string) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) AllowNoIndices(allownoindices bool) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) IgnoreUnavailable(ignoreunavailable bool) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) Resource(resource string) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) ErrorTrace(errortrace bool) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) FilterPath(filterpaths ...string) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) Human(human bool) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReloadSearchAnalyzers) Pretty(pretty bool) *ReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return nil
}
