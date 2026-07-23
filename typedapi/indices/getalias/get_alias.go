package getalias

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

	indexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetAlias struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name  string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetAlias func() *GetAlias

func NewGetAliasFunc(tp elastictransport.Interface) NewGetAlias {
	_ = "STUB: not implemented"
	return *new(NewGetAlias)
}

func New(tp elastictransport.Interface) *GetAlias { _ = "STUB: not implemented"; return nil }

func (r *GetAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAlias) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAlias) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetAlias) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetAlias) Header(key, value string) *GetAlias { _ = "STUB: not implemented"; return nil }

func (r *GetAlias) Name(name string) *GetAlias { _ = "STUB: not implemented"; return nil }

func (r *GetAlias) Index(index string) *GetAlias { _ = "STUB: not implemented"; return nil }

func (r *GetAlias) AllowNoIndices(allownoindices bool) *GetAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAlias) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAlias) IgnoreUnavailable(ignoreunavailable bool) *GetAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAlias) MasterTimeout(duration string) *GetAlias { _ = "STUB: not implemented"; return nil }

func (r *GetAlias) ErrorTrace(errortrace bool) *GetAlias { _ = "STUB: not implemented"; return nil }

func (r *GetAlias) FilterPath(filterpaths ...string) *GetAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAlias) Human(human bool) *GetAlias { _ = "STUB: not implemented"; return nil }

func (r *GetAlias) Pretty(pretty bool) *GetAlias { _ = "STUB: not implemented"; return nil }
