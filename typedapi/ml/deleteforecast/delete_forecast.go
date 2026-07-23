package deleteforecast

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	jobidMask = iota + 1

	forecastidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteForecast struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	jobid      string
	forecastid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteForecast func(jobid string) *DeleteForecast

func NewDeleteForecastFunc(tp elastictransport.Interface) NewDeleteForecast {
	_ = "STUB: not implemented"
	return *new(NewDeleteForecast)
}

func New(tp elastictransport.Interface) *DeleteForecast { _ = "STUB: not implemented"; return nil }

func (r *DeleteForecast) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteForecast) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteForecast) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteForecast) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteForecast) Header(key, value string) *DeleteForecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteForecast) _jobid(jobid string) *DeleteForecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteForecast) ForecastId(forecastid string) *DeleteForecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteForecast) AllowNoForecasts(allownoforecasts bool) *DeleteForecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteForecast) Timeout(duration string) *DeleteForecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteForecast) ErrorTrace(errortrace bool) *DeleteForecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteForecast) FilterPath(filterpaths ...string) *DeleteForecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteForecast) Human(human bool) *DeleteForecast { _ = "STUB: not implemented"; return nil }

func (r *DeleteForecast) Pretty(pretty bool) *DeleteForecast { _ = "STUB: not implemented"; return nil }
