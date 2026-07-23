package flamegraph

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Flamegraph struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      any
	deferred []func(request any) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewFlamegraph func() *Flamegraph

func NewFlamegraphFunc(tp elastictransport.Interface) NewFlamegraph {
	_ = "STUB: not implemented"
	return *new(NewFlamegraph)
}

func New(tp elastictransport.Interface) *Flamegraph { _ = "STUB: not implemented"; return nil }

func (r *Flamegraph) Raw(raw io.Reader) *Flamegraph { _ = "STUB: not implemented"; return nil }

func (r *Flamegraph) Request(req any) *Flamegraph { _ = "STUB: not implemented"; return nil }

func (r *Flamegraph) Conditions(conditions any) *Flamegraph { _ = "STUB: not implemented"; return nil }

func (r *Flamegraph) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Flamegraph) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Flamegraph) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *Flamegraph) Header(key, value string) *Flamegraph { _ = "STUB: not implemented"; return nil }

func (r *Flamegraph) ErrorTrace(errortrace bool) *Flamegraph { _ = "STUB: not implemented"; return nil }

func (r *Flamegraph) FilterPath(filterpaths ...string) *Flamegraph {
	_ = "STUB: not implemented"
	return nil
}

func (r *Flamegraph) Human(human bool) *Flamegraph { _ = "STUB: not implemented"; return nil }

func (r *Flamegraph) Pretty(pretty bool) *Flamegraph { _ = "STUB: not implemented"; return nil }
