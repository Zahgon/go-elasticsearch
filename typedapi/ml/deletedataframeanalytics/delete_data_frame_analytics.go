package deletedataframeanalytics

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

type DeleteDataFrameAnalytics struct {
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

type NewDeleteDataFrameAnalytics func(id string) *DeleteDataFrameAnalytics

func NewDeleteDataFrameAnalyticsFunc(tp elastictransport.Interface) NewDeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(NewDeleteDataFrameAnalytics)
}

func New(tp elastictransport.Interface) *DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataFrameAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataFrameAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataFrameAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDataFrameAnalytics) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteDataFrameAnalytics) Header(key, value string) *DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataFrameAnalytics) _id(id string) *DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataFrameAnalytics) Force(force bool) *DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataFrameAnalytics) Timeout(duration string) *DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataFrameAnalytics) ErrorTrace(errortrace bool) *DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataFrameAnalytics) FilterPath(filterpaths ...string) *DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataFrameAnalytics) Human(human bool) *DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDataFrameAnalytics) Pretty(pretty bool) *DeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return nil
}
