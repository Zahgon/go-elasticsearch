package forecast

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
	jobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Forecast struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewForecast func(jobid string) *Forecast

func NewForecastFunc(tp elastictransport.Interface) NewForecast {
	_ = "STUB: not implemented"
	return *new(NewForecast)
}

func New(tp elastictransport.Interface) *Forecast { _ = "STUB: not implemented"; return nil }

func (r *Forecast) Raw(raw io.Reader) *Forecast { _ = "STUB: not implemented"; return nil }

func (r *Forecast) Request(req *Request) *Forecast { _ = "STUB: not implemented"; return nil }

func (r *Forecast) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Forecast) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Forecast) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Forecast) Header(key, value string) *Forecast { _ = "STUB: not implemented"; return nil }

func (r *Forecast) _jobid(jobid string) *Forecast { _ = "STUB: not implemented"; return nil }

func (r *Forecast) ErrorTrace(errortrace bool) *Forecast { _ = "STUB: not implemented"; return nil }

func (r *Forecast) FilterPath(filterpaths ...string) *Forecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forecast) Human(human bool) *Forecast { _ = "STUB: not implemented"; return nil }

func (r *Forecast) Pretty(pretty bool) *Forecast { _ = "STUB: not implemented"; return nil }

func (r *Forecast) Duration(duration types.DurationVariant) *Forecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forecast) ExpiresIn(duration types.DurationVariant) *Forecast {
	_ = "STUB: not implemented"
	return nil
}

func (r *Forecast) MaxModelMemory(maxmodelmemory string) *Forecast {
	_ = "STUB: not implemented"
	return nil
}
