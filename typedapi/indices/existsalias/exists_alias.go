package existsalias

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

type ExistsAlias struct {
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

type NewExistsAlias func(name string) *ExistsAlias

func NewExistsAliasFunc(tp elastictransport.Interface) NewExistsAlias {
	_ = "STUB: not implemented"
	return *new(NewExistsAlias)
}

func New(tp elastictransport.Interface) *ExistsAlias { _ = "STUB: not implemented"; return nil }

func (r *ExistsAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsAlias) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsAlias) Do(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r ExistsAlias) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExistsAlias) Header(key, value string) *ExistsAlias { _ = "STUB: not implemented"; return nil }

func (r *ExistsAlias) _name(name string) *ExistsAlias { _ = "STUB: not implemented"; return nil }

func (r *ExistsAlias) Index(index string) *ExistsAlias { _ = "STUB: not implemented"; return nil }

func (r *ExistsAlias) AllowNoIndices(allownoindices bool) *ExistsAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsAlias) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ExistsAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsAlias) IgnoreUnavailable(ignoreunavailable bool) *ExistsAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsAlias) MasterTimeout(duration string) *ExistsAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsAlias) ErrorTrace(errortrace bool) *ExistsAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsAlias) FilterPath(filterpaths ...string) *ExistsAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsAlias) Human(human bool) *ExistsAlias { _ = "STUB: not implemented"; return nil }

func (r *ExistsAlias) Pretty(pretty bool) *ExistsAlias { _ = "STUB: not implemented"; return nil }
