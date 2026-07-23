package getjobstats

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
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetJobStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetJobStats func() *GetJobStats

func NewGetJobStatsFunc(tp elastictransport.Interface) NewGetJobStats {
	_ = "STUB: not implemented"
	return *new(NewGetJobStats)
}

func New(tp elastictransport.Interface) *GetJobStats { _ = "STUB: not implemented"; return nil }

func (r *GetJobStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetJobStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetJobStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetJobStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetJobStats) Header(key, value string) *GetJobStats { _ = "STUB: not implemented"; return nil }

func (r *GetJobStats) JobId(jobid string) *GetJobStats { _ = "STUB: not implemented"; return nil }

func (r *GetJobStats) AllowNoMatch(allownomatch bool) *GetJobStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetJobStats) ErrorTrace(errortrace bool) *GetJobStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetJobStats) FilterPath(filterpaths ...string) *GetJobStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetJobStats) Human(human bool) *GetJobStats { _ = "STUB: not implemented"; return nil }

func (r *GetJobStats) Pretty(pretty bool) *GetJobStats { _ = "STUB: not implemented"; return nil }
