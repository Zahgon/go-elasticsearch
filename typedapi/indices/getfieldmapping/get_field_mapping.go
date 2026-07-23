package getfieldmapping

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
	fieldsMask = iota + 1

	indexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetFieldMapping struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	fields string
	index  string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetFieldMapping func(fields string) *GetFieldMapping

func NewGetFieldMappingFunc(tp elastictransport.Interface) NewGetFieldMapping {
	_ = "STUB: not implemented"
	return *new(NewGetFieldMapping)
}

func New(tp elastictransport.Interface) *GetFieldMapping { _ = "STUB: not implemented"; return nil }

func (r *GetFieldMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFieldMapping) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetFieldMapping) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetFieldMapping) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetFieldMapping) Header(key, value string) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFieldMapping) _fields(fields string) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFieldMapping) Index(index string) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFieldMapping) AllowNoIndices(allownoindices bool) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFieldMapping) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFieldMapping) IgnoreUnavailable(ignoreunavailable bool) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFieldMapping) IncludeDefaults(includedefaults bool) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFieldMapping) ErrorTrace(errortrace bool) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFieldMapping) FilterPath(filterpaths ...string) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetFieldMapping) Human(human bool) *GetFieldMapping { _ = "STUB: not implemented"; return nil }

func (r *GetFieldMapping) Pretty(pretty bool) *GetFieldMapping {
	_ = "STUB: not implemented"
	return nil
}
