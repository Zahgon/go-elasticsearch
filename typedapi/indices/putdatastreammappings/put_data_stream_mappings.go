package putdatastreammappings

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/subobjects"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDataStreamMappings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutDataStreamMappings func(name string) *PutDataStreamMappings

func NewPutDataStreamMappingsFunc(tp elastictransport.Interface) NewPutDataStreamMappings {
	_ = "STUB: not implemented"
	return *new(NewPutDataStreamMappings)
}

func New(tp elastictransport.Interface) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Raw(raw io.Reader) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Request(req *Request) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataStreamMappings) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataStreamMappings) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutDataStreamMappings) Header(key, value string) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) _name(name string) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) DryRun(dryrun bool) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) MasterTimeout(duration string) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Timeout(duration string) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) ErrorTrace(errortrace bool) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) FilterPath(filterpaths ...string) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Human(human bool) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Pretty(pretty bool) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) AllField(allfield types.AllFieldVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) DataStreamTimestamp_(datastreamtimestamp_ types.DataStreamTimestampVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) DateDetection(datedetection bool) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Dynamic(dynamic dynamicmapping.DynamicMapping) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) DynamicDateFormats(dynamicdateformats ...string) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) DynamicTemplates(dynamictemplates []map[string]types.DynamicTemplate) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Enabled(enabled bool) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) FieldNames_(fieldnames_ types.FieldNamesFieldVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) IndexField(indexfield types.IndexFieldVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Meta_(metadata types.MetadataVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) NumericDetection(numericdetection bool) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Properties(properties map[string]types.Property) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) AddProperty(key string, value types.PropertyVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Routing_(routing_ types.RoutingFieldVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Runtime(runtime map[string]types.RuntimeField) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) AddRuntime(key string, value types.RuntimeFieldVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Size_(size_ types.SizeFieldVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Source_(source_ types.SourceFieldVariant) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataStreamMappings) Subobjects(subobjects subobjects.Subobjects) *PutDataStreamMappings {
	_ = "STUB: not implemented"
	return nil
}
