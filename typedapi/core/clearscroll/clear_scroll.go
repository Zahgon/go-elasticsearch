package clearscroll

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

type ClearScroll struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	scrollid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewClearScroll func() *ClearScroll

func NewClearScrollFunc(tp elastictransport.Interface) NewClearScroll {
	_ = "STUB: not implemented"
	return *new(NewClearScroll)
}

func New(tp elastictransport.Interface) *ClearScroll { _ = "STUB: not implemented"; return nil }

func (r *ClearScroll) Raw(raw io.Reader) *ClearScroll { _ = "STUB: not implemented"; return nil }

func (r *ClearScroll) Request(req *Request) *ClearScroll { _ = "STUB: not implemented"; return nil }

func (r *ClearScroll) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearScroll) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearScroll) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ClearScroll) Header(key, value string) *ClearScroll { _ = "STUB: not implemented"; return nil }

func (r *ClearScroll) ErrorTrace(errortrace bool) *ClearScroll {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearScroll) FilterPath(filterpaths ...string) *ClearScroll {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearScroll) Human(human bool) *ClearScroll { _ = "STUB: not implemented"; return nil }

func (r *ClearScroll) Pretty(pretty bool) *ClearScroll { _ = "STUB: not implemented"; return nil }

func (r *ClearScroll) ScrollId(scrollids ...string) *ClearScroll {
	_ = "STUB: not implemented"
	return nil
}
