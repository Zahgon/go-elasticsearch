package putmapping

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutMapping struct {
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

type NewPutMapping func(index string) *PutMapping

func NewPutMappingFunc(tp elastictransport.Interface) NewPutMapping {
	_ = "STUB: not implemented"
	return *new(NewPutMapping)
}

func New(tp elastictransport.Interface) *PutMapping { _ = "STUB: not implemented"; return nil }

func (r *PutMapping) Raw(raw io.Reader) *PutMapping { _ = "STUB: not implemented"; return nil }

func (r *PutMapping) Request(req *Request) *PutMapping { _ = "STUB: not implemented"; return nil }

func (r *PutMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutMapping) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutMapping) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutMapping) Header(key, value string) *PutMapping { _ = "STUB: not implemented"; return nil }

func (r *PutMapping) _index(index string) *PutMapping { _ = "STUB: not implemented"; return nil }

func (r *PutMapping) AllowNoIndices(allownoindices bool) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) IgnoreUnavailable(ignoreunavailable bool) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) MasterTimeout(duration string) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) Timeout(duration string) *PutMapping { _ = "STUB: not implemented"; return nil }

func (r *PutMapping) WriteIndexOnly(writeindexonly bool) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) ErrorTrace(errortrace bool) *PutMapping { _ = "STUB: not implemented"; return nil }

func (r *PutMapping) FilterPath(filterpaths ...string) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) Human(human bool) *PutMapping { _ = "STUB: not implemented"; return nil }

func (r *PutMapping) Pretty(pretty bool) *PutMapping { _ = "STUB: not implemented"; return nil }

func (r *PutMapping) DateDetection(datedetection bool) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) Dynamic(dynamic dynamicmapping.DynamicMapping) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) DynamicDateFormats(dynamicdateformats ...string) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) DynamicTemplates(dynamictemplates []map[string]types.DynamicTemplate) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) FieldNames_(fieldnames_ types.FieldNamesFieldVariant) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) Meta_(metadata types.MetadataVariant) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) NumericDetection(numericdetection bool) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) Properties(properties map[string]types.Property) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) AddProperty(key string, value types.PropertyVariant) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) Routing_(routing_ types.RoutingFieldVariant) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) Runtime(runtimefields types.RuntimeFieldsVariant) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMapping) Source_(source_ types.SourceFieldVariant) *PutMapping {
	_ = "STUB: not implemented"
	return nil
}
