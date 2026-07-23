package getstats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetStats func() *GetStats

func NewGetStatsFunc(tp elastictransport.Interface) NewGetStats {
	_ = "STUB: not implemented"
	return *new(NewGetStats)
}

func New(tp elastictransport.Interface) *GetStats { _ = "STUB: not implemented"; return nil }

func (r *GetStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetStats) Header(key, value string) *GetStats { _ = "STUB: not implemented"; return nil }

func (r *GetStats) ErrorTrace(errortrace bool) *GetStats { _ = "STUB: not implemented"; return nil }

func (r *GetStats) FilterPath(filterpaths ...string) *GetStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetStats) Human(human bool) *GetStats { _ = "STUB: not implemented"; return nil }

func (r *GetStats) Pretty(pretty bool) *GetStats { _ = "STUB: not implemented"; return nil }
