package healthreport

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	featureMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type HealthReport struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	feature string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewHealthReport func() *HealthReport

func NewHealthReportFunc(tp elastictransport.Interface) NewHealthReport {
	_ = "STUB: not implemented"
	return *new(NewHealthReport)
}

func New(tp elastictransport.Interface) *HealthReport { _ = "STUB: not implemented"; return nil }

func (r *HealthReport) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HealthReport) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HealthReport) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HealthReport) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *HealthReport) Header(key, value string) *HealthReport {
	_ = "STUB: not implemented"
	return nil
}

func (r *HealthReport) Feature(features ...string) *HealthReport {
	_ = "STUB: not implemented"
	return nil
}

func (r *HealthReport) Timeout(duration string) *HealthReport {
	_ = "STUB: not implemented"
	return nil
}

func (r *HealthReport) Verbose(verbose bool) *HealthReport { _ = "STUB: not implemented"; return nil }

func (r *HealthReport) Size(size int) *HealthReport { _ = "STUB: not implemented"; return nil }

func (r *HealthReport) ErrorTrace(errortrace bool) *HealthReport {
	_ = "STUB: not implemented"
	return nil
}

func (r *HealthReport) FilterPath(filterpaths ...string) *HealthReport {
	_ = "STUB: not implemented"
	return nil
}

func (r *HealthReport) Human(human bool) *HealthReport { _ = "STUB: not implemented"; return nil }

func (r *HealthReport) Pretty(pretty bool) *HealthReport { _ = "STUB: not implemented"; return nil }
