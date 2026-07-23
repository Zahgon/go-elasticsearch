package putbehavioralanalytics

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

type PutBehavioralAnalytics struct {
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

type NewPutBehavioralAnalytics func(name string) *PutBehavioralAnalytics

func NewPutBehavioralAnalyticsFunc(tp elastictransport.Interface) NewPutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return *new(NewPutBehavioralAnalytics)
}

func New(tp elastictransport.Interface) *PutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutBehavioralAnalytics) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutBehavioralAnalytics) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutBehavioralAnalytics) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutBehavioralAnalytics) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PutBehavioralAnalytics) Header(key, value string) *PutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutBehavioralAnalytics) _name(name string) *PutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutBehavioralAnalytics) ErrorTrace(errortrace bool) *PutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutBehavioralAnalytics) FilterPath(filterpaths ...string) *PutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutBehavioralAnalytics) Human(human bool) *PutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutBehavioralAnalytics) Pretty(pretty bool) *PutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return nil
}
