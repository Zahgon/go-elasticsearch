package getmemorystats

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

type GetMemoryStats struct {
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

type NewGetMemoryStats func() *GetMemoryStats

func NewGetMemoryStatsFunc(tp elastictransport.Interface) NewGetMemoryStats {
	_ = "STUB: not implemented"
	return *new(NewGetMemoryStats)
}

func New(tp elastictransport.Interface) *GetMemoryStats { _ = "STUB: not implemented"; return nil }

func (r *GetMemoryStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetMemoryStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetMemoryStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetMemoryStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetMemoryStats) Header(key, value string) *GetMemoryStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMemoryStats) NodeId(nodeid string) *GetMemoryStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMemoryStats) MasterTimeout(duration string) *GetMemoryStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMemoryStats) Timeout(duration string) *GetMemoryStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMemoryStats) ErrorTrace(errortrace bool) *GetMemoryStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMemoryStats) FilterPath(filterpaths ...string) *GetMemoryStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMemoryStats) Human(human bool) *GetMemoryStats { _ = "STUB: not implemented"; return nil }

func (r *GetMemoryStats) Pretty(pretty bool) *GetMemoryStats { _ = "STUB: not implemented"; return nil }
