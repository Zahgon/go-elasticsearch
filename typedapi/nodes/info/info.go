package info

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

type Info struct {
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

type NewInfo func() *Info

func NewInfoFunc(tp elastictransport.Interface) NewInfo {
	_ = "STUB: not implemented"
	return *new(NewInfo)
}

func New(tp elastictransport.Interface) *Info { _ = "STUB: not implemented"; return nil }

func (r *Info) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Info) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Info) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Info) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Info) Header(key, value string) *Info { _ = "STUB: not implemented"; return nil }

func (r *Info) NodeId(nodeid string) *Info { _ = "STUB: not implemented"; return nil }

func (r *Info) Metric(metric string) *Info { _ = "STUB: not implemented"; return nil }

func (r *Info) FlatSettings(flatsettings bool) *Info { _ = "STUB: not implemented"; return nil }

func (r *Info) Timeout(duration string) *Info { _ = "STUB: not implemented"; return nil }

func (r *Info) ErrorTrace(errortrace bool) *Info { _ = "STUB: not implemented"; return nil }

func (r *Info) FilterPath(filterpaths ...string) *Info { _ = "STUB: not implemented"; return nil }

func (r *Info) Human(human bool) *Info { _ = "STUB: not implemented"; return nil }

func (r *Info) Pretty(pretty bool) *Info { _ = "STUB: not implemented"; return nil }
