package putdataframeanalytics

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

type PutDataFrameAnalytics struct {
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

type NewPutDataFrameAnalytics func(id string) *PutDataFrameAnalytics

func NewPutDataFrameAnalyticsFunc(tp elastictransport.Interface) NewPutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(NewPutDataFrameAnalytics)
}

func New(tp elastictransport.Interface) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Raw(raw io.Reader) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Request(req *Request) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutDataFrameAnalytics) Header(key, value string) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) _id(id string) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) ErrorTrace(errortrace bool) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) FilterPath(filterpaths ...string) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Human(human bool) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Pretty(pretty bool) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) AllowLazyStart(allowlazystart bool) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Analysis(analysis types.DataframeAnalysisContainerVariant) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) AnalyzedFields(analyzedfields types.DataframeAnalysisAnalyzedFieldsVariant) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Description(description string) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Dest(dest types.DataframeAnalyticsDestinationVariant) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Headers(httpheaders types.HttpHeadersVariant) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) MaxNumThreads(maxnumthreads int) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Meta_(metadata types.MetadataVariant) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) ModelMemoryLimit(modelmemorylimit string) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Source(source types.DataframeAnalyticsSourceVariant) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDataFrameAnalytics) Version(versionstring string) *PutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}
