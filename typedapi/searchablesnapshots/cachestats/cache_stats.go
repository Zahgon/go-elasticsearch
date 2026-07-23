package cachestats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nodeidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CacheStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	nodeid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCacheStats func() *CacheStats

func NewCacheStatsFunc(tp elastictransport.Interface) NewCacheStats {
	_ = "STUB: not implemented"
	return *new(NewCacheStats)
}

func New(tp elastictransport.Interface) *CacheStats { _ = "STUB: not implemented"; return nil }

func (r *CacheStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CacheStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CacheStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CacheStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *CacheStats) Header(key, value string) *CacheStats { _ = "STUB: not implemented"; return nil }

func (r *CacheStats) NodeId(nodeid string) *CacheStats { _ = "STUB: not implemented"; return nil }

func (r *CacheStats) ErrorTrace(errortrace bool) *CacheStats { _ = "STUB: not implemented"; return nil }

func (r *CacheStats) FilterPath(filterpaths ...string) *CacheStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *CacheStats) Human(human bool) *CacheStats { _ = "STUB: not implemented"; return nil }

func (r *CacheStats) Pretty(pretty bool) *CacheStats { _ = "STUB: not implemented"; return nil }
