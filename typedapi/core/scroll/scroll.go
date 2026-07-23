package scroll

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

type Scroll struct {
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

type NewScroll func() *Scroll

func NewScrollFunc(tp elastictransport.Interface) NewScroll {
	_ = "STUB: not implemented"
	return *new(NewScroll)
}

func New(tp elastictransport.Interface) *Scroll { _ = "STUB: not implemented"; return nil }

func (r *Scroll) Raw(raw io.Reader) *Scroll { _ = "STUB: not implemented"; return nil }

func (r *Scroll) Request(req *Request) *Scroll { _ = "STUB: not implemented"; return nil }

func (r *Scroll) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Scroll) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Scroll) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Scroll) Header(key, value string) *Scroll { _ = "STUB: not implemented"; return nil }

func (r *Scroll) RestTotalHitsAsInt(resttotalhitsasint bool) *Scroll {
	_ = "STUB: not implemented"
	return nil
}

func (r *Scroll) ErrorTrace(errortrace bool) *Scroll { _ = "STUB: not implemented"; return nil }

func (r *Scroll) FilterPath(filterpaths ...string) *Scroll { _ = "STUB: not implemented"; return nil }

func (r *Scroll) Human(human bool) *Scroll { _ = "STUB: not implemented"; return nil }

func (r *Scroll) Pretty(pretty bool) *Scroll { _ = "STUB: not implemented"; return nil }

func (r *Scroll) Scroll(duration types.DurationVariant) *Scroll {
	_ = "STUB: not implemented"
	return nil
}

func (r *Scroll) ScrollId(scrollid string) *Scroll { _ = "STUB: not implemented"; return nil }
