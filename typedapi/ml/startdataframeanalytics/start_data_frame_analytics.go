package startdataframeanalytics

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

type StartDataFrameAnalytics struct {
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

type NewStartDataFrameAnalytics func(id string) *StartDataFrameAnalytics

func NewStartDataFrameAnalyticsFunc(tp elastictransport.Interface) NewStartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(NewStartDataFrameAnalytics)
}

func New(tp elastictransport.Interface) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDataFrameAnalytics) Raw(raw io.Reader) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDataFrameAnalytics) Request(req *Request) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *StartDataFrameAnalytics) Header(key, value string) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDataFrameAnalytics) _id(id string) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDataFrameAnalytics) ErrorTrace(errortrace bool) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDataFrameAnalytics) FilterPath(filterpaths ...string) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDataFrameAnalytics) Human(human bool) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDataFrameAnalytics) Pretty(pretty bool) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDataFrameAnalytics) Timeout(duration types.DurationVariant) *StartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}
