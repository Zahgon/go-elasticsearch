package reroute

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

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Reroute struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewReroute func() *Reroute

func NewRerouteFunc(tp elastictransport.Interface) NewReroute {
	_ = "STUB: not implemented"
	return *new(NewReroute)
}

func New(tp elastictransport.Interface) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) Raw(raw io.Reader) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) Request(req *Request) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Reroute) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Reroute) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reroute) Header(key, value string) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) DryRun(dryrun bool) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) Explain(explain bool) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) Metric(metrics ...string) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) RetryFailed(retryfailed bool) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) MasterTimeout(duration string) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) Timeout(duration string) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) ErrorTrace(errortrace bool) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) FilterPath(filterpaths ...string) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) Human(human bool) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) Pretty(pretty bool) *Reroute { _ = "STUB: not implemented"; return nil }

func (r *Reroute) Commands(commands ...types.CommandVariant) *Reroute {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reroute) CommandsValues(commandsvalues []types.Command) *Reroute {
	_ = "STUB: not implemented"
	return nil
}
