package previewdataframeanalytics

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

type PreviewDataFrameAnalytics struct {
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

type NewPreviewDataFrameAnalytics func() *PreviewDataFrameAnalytics

func NewPreviewDataFrameAnalyticsFunc(tp elastictransport.Interface) NewPreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(NewPreviewDataFrameAnalytics)
}

func New(tp elastictransport.Interface) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDataFrameAnalytics) Raw(raw io.Reader) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDataFrameAnalytics) Request(req *Request) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PreviewDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PreviewDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PreviewDataFrameAnalytics) Header(key, value string) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDataFrameAnalytics) Id(id string) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDataFrameAnalytics) ErrorTrace(errortrace bool) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDataFrameAnalytics) FilterPath(filterpaths ...string) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDataFrameAnalytics) Human(human bool) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDataFrameAnalytics) Pretty(pretty bool) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDataFrameAnalytics) Config(config types.DataframePreviewConfigVariant) *PreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}
