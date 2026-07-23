package putelasticsearch

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/elasticsearchservicetype"
)

const (
	tasktypeMask = iota + 1

	elasticsearchinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutElasticsearch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                 string
	elasticsearchinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutElasticsearch func(tasktype, elasticsearchinferenceid string) *PutElasticsearch

func NewPutElasticsearchFunc(tp elastictransport.Interface) NewPutElasticsearch {
	_ = "STUB: not implemented"
	return *new(NewPutElasticsearch)
}

func New(tp elastictransport.Interface) *PutElasticsearch { _ = "STUB: not implemented"; return nil }

func (r *PutElasticsearch) Raw(raw io.Reader) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) Request(req *Request) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutElasticsearch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutElasticsearch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutElasticsearch) Header(key, value string) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) _tasktype(tasktype string) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) _elasticsearchinferenceid(elasticsearchinferenceid string) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) Timeout(duration string) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) ErrorTrace(errortrace bool) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) FilterPath(filterpaths ...string) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) Human(human bool) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) Pretty(pretty bool) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) Service(service elasticsearchservicetype.ElasticsearchServiceType) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) ServiceSettings(servicesettings types.ElasticsearchServiceSettingsVariant) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElasticsearch) TaskSettings(tasksettings types.ElasticsearchTaskSettingsVariant) *PutElasticsearch {
	_ = "STUB: not implemented"
	return nil
}
