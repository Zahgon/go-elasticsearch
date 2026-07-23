package fieldcaps

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type FieldCaps struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewFieldCaps func() *FieldCaps

func NewFieldCapsFunc(tp elastictransport.Interface) NewFieldCaps {
	_ = "STUB: not implemented"
	return *new(NewFieldCaps)
}

func New(tp elastictransport.Interface) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) Raw(raw io.Reader) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) Request(req *Request) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FieldCaps) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FieldCaps) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *FieldCaps) Header(key, value string) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) Index(index string) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) AllowNoIndices(allownoindices bool) *FieldCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldCaps) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *FieldCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldCaps) IgnoreUnavailable(ignoreunavailable bool) *FieldCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldCaps) IncludeUnmapped(includeunmapped bool) *FieldCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldCaps) Filters(filters ...string) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) Types(types ...string) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) IncludeEmptyFields(includeemptyfields bool) *FieldCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldCaps) ErrorTrace(errortrace bool) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) FilterPath(filterpaths ...string) *FieldCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldCaps) Human(human bool) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) Pretty(pretty bool) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) Fields(fields ...string) *FieldCaps { _ = "STUB: not implemented"; return nil }

func (r *FieldCaps) IndexFilter(indexfilter types.QueryVariant) *FieldCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldCaps) ProjectRouting(projectrouting string) *FieldCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *FieldCaps) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *FieldCaps {
	_ = "STUB: not implemented"
	return nil
}
