package getbehavioralanalytics

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

type GetBehavioralAnalytics struct {
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

type NewGetBehavioralAnalytics func() *GetBehavioralAnalytics

func NewGetBehavioralAnalyticsFunc(tp elastictransport.Interface) NewGetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return *new(NewGetBehavioralAnalytics)
}

func New(tp elastictransport.Interface) *GetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBehavioralAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBehavioralAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBehavioralAnalytics) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetBehavioralAnalytics) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetBehavioralAnalytics) Header(key, value string) *GetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBehavioralAnalytics) Name(name string) *GetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBehavioralAnalytics) ErrorTrace(errortrace bool) *GetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBehavioralAnalytics) FilterPath(filterpaths ...string) *GetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBehavioralAnalytics) Human(human bool) *GetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBehavioralAnalytics) Pretty(pretty bool) *GetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}
