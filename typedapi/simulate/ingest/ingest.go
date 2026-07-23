package ingest

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/mergetype"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Ingest struct {
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

type NewIngest func() *Ingest

func NewIngestFunc(tp elastictransport.Interface) NewIngest {
	_ = "STUB: not implemented"
	return *new(NewIngest)
}

func New(tp elastictransport.Interface) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) Raw(raw io.Reader) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) Request(req *Request) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Ingest) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Ingest) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Ingest) Header(key, value string) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) Index(index string) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) Pipeline(pipelinename string) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) MergeType(mergetype mergetype.MergeType) *Ingest {
	_ = "STUB: not implemented"
	return nil
}

func (r *Ingest) ErrorTrace(errortrace bool) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) FilterPath(filterpaths ...string) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) Human(human bool) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) Pretty(pretty bool) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) ComponentTemplateSubstitutions(componenttemplatesubstitutions map[string]types.ComponentTemplateNode) *Ingest {
	_ = "STUB: not implemented"
	return nil
}

func (r *Ingest) AddComponentTemplateSubstitution(key string, value types.ComponentTemplateNodeVariant) *Ingest {
	_ = "STUB: not implemented"
	return nil
}

func (r *Ingest) Docs(docs ...types.DocumentVariant) *Ingest { _ = "STUB: not implemented"; return nil }

func (r *Ingest) DocsValues(docsvalues []types.Document) *Ingest {
	_ = "STUB: not implemented"
	return nil
}

func (r *Ingest) IndexTemplateSubstitutions(indextemplatesubstitutions map[string]types.IndexTemplate) *Ingest {
	_ = "STUB: not implemented"
	return nil
}

func (r *Ingest) AddIndexTemplateSubstitution(key string, value types.IndexTemplateVariant) *Ingest {
	_ = "STUB: not implemented"
	return nil
}

func (r *Ingest) MappingAddition(mappingaddition types.TypeMappingVariant) *Ingest {
	_ = "STUB: not implemented"
	return nil
}

func (r *Ingest) PipelineSubstitutions(pipelinesubstitutions map[string]types.IngestPipeline) *Ingest {
	_ = "STUB: not implemented"
	return nil
}

func (r *Ingest) AddPipelineSubstitution(key string, value types.IngestPipelineVariant) *Ingest {
	_ = "STUB: not implemented"
	return nil
}
