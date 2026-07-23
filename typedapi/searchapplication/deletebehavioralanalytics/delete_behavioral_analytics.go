package deletebehavioralanalytics

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteBehavioralAnalytics struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteBehavioralAnalytics func(name string) *DeleteBehavioralAnalytics

func NewDeleteBehavioralAnalyticsFunc(tp elastictransport.Interface) NewDeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return *new(NewDeleteBehavioralAnalytics)
}

func New(tp elastictransport.Interface) *DeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteBehavioralAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteBehavioralAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteBehavioralAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteBehavioralAnalytics) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteBehavioralAnalytics) Header(key, value string) *DeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteBehavioralAnalytics) _name(name string) *DeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteBehavioralAnalytics) ErrorTrace(errortrace bool) *DeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteBehavioralAnalytics) FilterPath(filterpaths ...string) *DeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteBehavioralAnalytics) Human(human bool) *DeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteBehavioralAnalytics) Pretty(pretty bool) *DeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}
