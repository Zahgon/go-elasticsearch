package exists

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

type Exists struct {
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

type NewExists func(index string) *Exists

func NewExistsFunc(tp elastictransport.Interface) NewExists {
	_ = "STUB: not implemented"
	return *new(NewExists)
}

func New(tp elastictransport.Interface) *Exists { _ = "STUB: not implemented"; return nil }

func (r *Exists) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Exists) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Exists) Do(ctx context.Context) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r Exists) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Exists) Header(key, value string) *Exists { _ = "STUB: not implemented"; return nil }

func (r *Exists) _index(index string) *Exists { _ = "STUB: not implemented"; return nil }

func (r *Exists) AllowNoIndices(allownoindices bool) *Exists { _ = "STUB: not implemented"; return nil }

func (r *Exists) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Exists {
	_ = "STUB: not implemented"
	return nil
}

func (r *Exists) FlatSettings(flatsettings bool) *Exists { _ = "STUB: not implemented"; return nil }

func (r *Exists) IgnoreUnavailable(ignoreunavailable bool) *Exists {
	_ = "STUB: not implemented"
	return nil
}

func (r *Exists) IncludeDefaults(includedefaults bool) *Exists {
	_ = "STUB: not implemented"
	return nil
}

func (r *Exists) Local(local bool) *Exists { _ = "STUB: not implemented"; return nil }

func (r *Exists) ErrorTrace(errortrace bool) *Exists { _ = "STUB: not implemented"; return nil }

func (r *Exists) FilterPath(filterpaths ...string) *Exists { _ = "STUB: not implemented"; return nil }

func (r *Exists) Human(human bool) *Exists { _ = "STUB: not implemented"; return nil }

func (r *Exists) Pretty(pretty bool) *Exists { _ = "STUB: not implemented"; return nil }
