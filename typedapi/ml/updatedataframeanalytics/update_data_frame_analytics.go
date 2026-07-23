package updatedataframeanalytics

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateDataFrameAnalytics struct {
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

type NewUpdateDataFrameAnalytics func(id string) *UpdateDataFrameAnalytics

func NewUpdateDataFrameAnalyticsFunc(tp elastictransport.Interface) NewUpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(NewUpdateDataFrameAnalytics)
}

func New(tp elastictransport.Interface) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) Raw(raw io.Reader) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) Request(req *Request) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateDataFrameAnalytics) Header(key, value string) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) _id(id string) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) ErrorTrace(errortrace bool) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) FilterPath(filterpaths ...string) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) Human(human bool) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) Pretty(pretty bool) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) AllowLazyStart(allowlazystart bool) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) Description(description string) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) MaxNumThreads(maxnumthreads int) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDataFrameAnalytics) ModelMemoryLimit(modelmemorylimit string) *UpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}
