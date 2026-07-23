package stacktraces

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

type Stacktraces struct {
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

type NewStacktraces func() *Stacktraces

func NewStacktracesFunc(tp elastictransport.Interface) NewStacktraces {
	_ = "STUB: not implemented"
	return *new(NewStacktraces)
}

func New(tp elastictransport.Interface) *Stacktraces { _ = "STUB: not implemented"; return nil }

func (r *Stacktraces) Raw(raw io.Reader) *Stacktraces { _ = "STUB: not implemented"; return nil }

func (r *Stacktraces) Request(req any) *Stacktraces { _ = "STUB: not implemented"; return nil }

func (r *Stacktraces) Conditions(conditions any) *Stacktraces {
	_ = "STUB: not implemented"
	return nil
}

func (r *Stacktraces) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stacktraces) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stacktraces) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *Stacktraces) Header(key, value string) *Stacktraces { _ = "STUB: not implemented"; return nil }

func (r *Stacktraces) ErrorTrace(errortrace bool) *Stacktraces {
	_ = "STUB: not implemented"
	return nil
}

func (r *Stacktraces) FilterPath(filterpaths ...string) *Stacktraces {
	_ = "STUB: not implemented"
	return nil
}

func (r *Stacktraces) Human(human bool) *Stacktraces { _ = "STUB: not implemented"; return nil }

func (r *Stacktraces) Pretty(pretty bool) *Stacktraces { _ = "STUB: not implemented"; return nil }
