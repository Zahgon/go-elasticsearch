package stopdataframeanalytics

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

type StopDataFrameAnalytics struct {
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

type NewStopDataFrameAnalytics func(id string) *StopDataFrameAnalytics

func NewStopDataFrameAnalyticsFunc(tp elastictransport.Interface) NewStopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(NewStopDataFrameAnalytics)
}

func New(tp elastictransport.Interface) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) Raw(raw io.Reader) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) Request(req *Request) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *StopDataFrameAnalytics) Header(key, value string) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) _id(id string) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) ErrorTrace(errortrace bool) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) FilterPath(filterpaths ...string) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) Human(human bool) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) Pretty(pretty bool) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) AllowNoMatch(allownomatch bool) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) Force(force bool) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDataFrameAnalytics) Timeout(duration types.DurationVariant) *StopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}
