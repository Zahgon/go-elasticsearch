package getdataframeanalyticsstats

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

type GetDataFrameAnalyticsStats struct {
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

type NewGetDataFrameAnalyticsStats func() *GetDataFrameAnalyticsStats

func NewGetDataFrameAnalyticsStatsFunc(tp elastictransport.Interface) NewGetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return *new(NewGetDataFrameAnalyticsStats)
}

func New(tp elastictransport.Interface) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataFrameAnalyticsStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataFrameAnalyticsStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataFrameAnalyticsStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDataFrameAnalyticsStats) Header(key, value string) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) Id(id string) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) AllowNoMatch(allownomatch bool) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) From(from int) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) Size(size int) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) Verbose(verbose bool) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) ErrorTrace(errortrace bool) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) FilterPath(filterpaths ...string) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) Human(human bool) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataFrameAnalyticsStats) Pretty(pretty bool) *GetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return nil
}
