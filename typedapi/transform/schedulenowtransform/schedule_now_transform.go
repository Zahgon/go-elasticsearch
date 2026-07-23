package schedulenowtransform

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	transformidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ScheduleNowTransform struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	transformid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewScheduleNowTransform func(transformid string) *ScheduleNowTransform

func NewScheduleNowTransformFunc(tp elastictransport.Interface) NewScheduleNowTransform {
	_ = "STUB: not implemented"
	return *new(NewScheduleNowTransform)
}

func New(tp elastictransport.Interface) *ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScheduleNowTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ScheduleNowTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ScheduleNowTransform) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ScheduleNowTransform) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ScheduleNowTransform) Header(key, value string) *ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScheduleNowTransform) _transformid(transformid string) *ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScheduleNowTransform) Timeout(duration string) *ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScheduleNowTransform) Defer(defer_ bool) *ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScheduleNowTransform) ErrorTrace(errortrace bool) *ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScheduleNowTransform) FilterPath(filterpaths ...string) *ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScheduleNowTransform) Human(human bool) *ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ScheduleNowTransform) Pretty(pretty bool) *ScheduleNowTransform {
	_ = "STUB: not implemented"
	return nil
}
