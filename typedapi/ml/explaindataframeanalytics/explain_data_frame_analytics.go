package explaindataframeanalytics

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

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExplainDataFrameAnalytics struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewExplainDataFrameAnalytics func() *ExplainDataFrameAnalytics

func NewExplainDataFrameAnalyticsFunc(tp elastictransport.Interface) NewExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(NewExplainDataFrameAnalytics)
}

func New(tp elastictransport.Interface) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) Raw(raw io.Reader) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) Request(req *Request) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExplainDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExplainDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ExplainDataFrameAnalytics) Header(key, value string) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) Id(id string) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) ErrorTrace(errortrace bool) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) FilterPath(filterpaths ...string) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) Human(human bool) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) Pretty(pretty bool) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) AllowLazyStart(allowlazystart bool) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) Analysis(analysis types.DataframeAnalysisContainerVariant) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) AnalyzedFields(analyzedfields types.DataframeAnalysisAnalyzedFieldsVariant) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) Description(description string) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) Dest(dest types.DataframeAnalyticsDestinationVariant) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) MaxNumThreads(maxnumthreads int) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) ModelMemoryLimit(modelmemorylimit string) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExplainDataFrameAnalytics) Source(source types.DataframeAnalyticsSourceVariant) *ExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}
