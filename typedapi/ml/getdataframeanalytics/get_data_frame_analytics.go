package getdataframeanalytics

import (
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

type GetDataFrameAnalytics struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetDataFrameAnalytics func() *GetDataFrameAnalytics

func NewGetDataFrameAnalyticsFunc(tp elastictransport.Interface) NewGetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(NewGetDataFrameAnalytics)
}

func New(tp elastictransport.Interface) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataFrameAnalytics) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDataFrameAnalytics) Header(key, value string) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) Id(id string) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) AllowNoMatch(allownomatch bool) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) From(from int) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) Size(size int) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) ExcludeGenerated(excludegenerated bool) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) ErrorTrace(errortrace bool) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) FilterPath(filterpaths ...string) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) Human(human bool) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalytics) Pretty(pretty bool) *GetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}
