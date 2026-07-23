package getmapping

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

type GetMapping struct {
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

type NewGetMapping func() *GetMapping

func NewGetMappingFunc(tp elastictransport.Interface) NewGetMapping {
	_ = "STUB: not implemented"
	return *new(NewGetMapping)
}

func New(tp elastictransport.Interface) *GetMapping { _ = "STUB: not implemented"; return nil }

func (r *GetMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetMapping) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetMapping) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetMapping) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetMapping) Header(key, value string) *GetMapping { _ = "STUB: not implemented"; return nil }

func (r *GetMapping) Index(index string) *GetMapping { _ = "STUB: not implemented"; return nil }

func (r *GetMapping) AllowNoIndices(allownoindices bool) *GetMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMapping) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMapping) IgnoreUnavailable(ignoreunavailable bool) *GetMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMapping) Local(local bool) *GetMapping { _ = "STUB: not implemented"; return nil }

func (r *GetMapping) MasterTimeout(duration string) *GetMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMapping) ErrorTrace(errortrace bool) *GetMapping { _ = "STUB: not implemented"; return nil }

func (r *GetMapping) FilterPath(filterpaths ...string) *GetMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMapping) Human(human bool) *GetMapping { _ = "STUB: not implemented"; return nil }

func (r *GetMapping) Pretty(pretty bool) *GetMapping { _ = "STUB: not implemented"; return nil }
