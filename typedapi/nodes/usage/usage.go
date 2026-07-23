package usage

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

	metricMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Usage struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	nodeid string
	metric string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUsage func() *Usage

func NewUsageFunc(tp elastictransport.Interface) NewUsage {
	_ = "STUB: not implemented"
	return *new(NewUsage)
}

func New(tp elastictransport.Interface) *Usage { _ = "STUB: not implemented"; return nil }

func (r *Usage) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Usage) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Usage) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Usage) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Usage) Header(key, value string) *Usage { _ = "STUB: not implemented"; return nil }

func (r *Usage) NodeId(nodeid string) *Usage { _ = "STUB: not implemented"; return nil }

func (r *Usage) Metric(metric string) *Usage { _ = "STUB: not implemented"; return nil }

func (r *Usage) Timeout(duration string) *Usage { _ = "STUB: not implemented"; return nil }

func (r *Usage) ErrorTrace(errortrace bool) *Usage { _ = "STUB: not implemented"; return nil }

func (r *Usage) FilterPath(filterpaths ...string) *Usage { _ = "STUB: not implemented"; return nil }

func (r *Usage) Human(human bool) *Usage { _ = "STUB: not implemented"; return nil }

func (r *Usage) Pretty(pretty bool) *Usage { _ = "STUB: not implemented"; return nil }
