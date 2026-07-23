package getdatalifecyclestats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetDataLifecycleStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetDataLifecycleStats func() *GetDataLifecycleStats

func NewGetDataLifecycleStatsFunc(tp elastictransport.Interface) NewGetDataLifecycleStats {
	_ = "STUB: not implemented"
	return *new(NewGetDataLifecycleStats)
}

func New(tp elastictransport.Interface) *GetDataLifecycleStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycleStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataLifecycleStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataLifecycleStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDataLifecycleStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDataLifecycleStats) Header(key, value string) *GetDataLifecycleStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycleStats) ErrorTrace(errortrace bool) *GetDataLifecycleStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycleStats) FilterPath(filterpaths ...string) *GetDataLifecycleStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycleStats) Human(human bool) *GetDataLifecycleStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDataLifecycleStats) Pretty(pretty bool) *GetDataLifecycleStats {
	_ = "STUB: not implemented"
	return nil
}
