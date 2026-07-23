package estimatemodelmemory

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type EstimateModelMemory struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewEstimateModelMemory func() *EstimateModelMemory

func NewEstimateModelMemoryFunc(tp elastictransport.Interface) NewEstimateModelMemory {
	_ = "STUB: not implemented"
	return *new(NewEstimateModelMemory)
}

func New(tp elastictransport.Interface) *EstimateModelMemory { _ = "STUB: not implemented"; return nil }

func (r *EstimateModelMemory) Raw(raw io.Reader) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) Request(req *Request) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EstimateModelMemory) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EstimateModelMemory) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *EstimateModelMemory) Header(key, value string) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) ErrorTrace(errortrace bool) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) FilterPath(filterpaths ...string) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) Human(human bool) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) Pretty(pretty bool) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) AnalysisConfig(analysisconfig types.AnalysisConfigVariant) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) MaxBucketCardinality(maxbucketcardinality map[string]int64) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) AddMaxBucketCardinality(key string, value int64) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) OverallCardinality(overallcardinality map[string]int64) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}

func (r *EstimateModelMemory) AddOverallCardinality(key string, value int64) *EstimateModelMemory {
	_ = "STUB: not implemented"
	return nil
}
