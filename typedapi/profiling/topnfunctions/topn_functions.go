package topnfunctions

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

type TopnFunctions struct {
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

type NewTopnFunctions func() *TopnFunctions

func NewTopnFunctionsFunc(tp elastictransport.Interface) NewTopnFunctions {
	_ = "STUB: not implemented"
	return *new(NewTopnFunctions)
}

func New(tp elastictransport.Interface) *TopnFunctions { _ = "STUB: not implemented"; return nil }

func (r *TopnFunctions) Raw(raw io.Reader) *TopnFunctions { _ = "STUB: not implemented"; return nil }

func (r *TopnFunctions) Request(req any) *TopnFunctions { _ = "STUB: not implemented"; return nil }

func (r *TopnFunctions) Conditions(conditions any) *TopnFunctions {
	_ = "STUB: not implemented"
	return nil
}

func (r *TopnFunctions) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r TopnFunctions) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r TopnFunctions) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *TopnFunctions) Header(key, value string) *TopnFunctions {
	_ = "STUB: not implemented"
	return nil
}

func (r *TopnFunctions) ErrorTrace(errortrace bool) *TopnFunctions {
	_ = "STUB: not implemented"
	return nil
}

func (r *TopnFunctions) FilterPath(filterpaths ...string) *TopnFunctions {
	_ = "STUB: not implemented"
	return nil
}

func (r *TopnFunctions) Human(human bool) *TopnFunctions { _ = "STUB: not implemented"; return nil }

func (r *TopnFunctions) Pretty(pretty bool) *TopnFunctions { _ = "STUB: not implemented"; return nil }
