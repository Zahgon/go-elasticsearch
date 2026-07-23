package stats

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

type Stats struct {
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

type NewStats func() *Stats

func NewStatsFunc(tp elastictransport.Interface) NewStats {
	_ = "STUB: not implemented"
	return *new(NewStats)
}

func New(tp elastictransport.Interface) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Stats) Header(key, value string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) NodeId(nodeid string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) IncludeRemotes(includeremotes bool) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) Timeout(duration string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) ErrorTrace(errortrace bool) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) FilterPath(filterpaths ...string) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) Human(human bool) *Stats { _ = "STUB: not implemented"; return nil }

func (r *Stats) Pretty(pretty bool) *Stats { _ = "STUB: not implemented"; return nil }
