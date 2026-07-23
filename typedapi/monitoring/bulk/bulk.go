package bulk

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

type Bulk struct {
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

type NewBulk func() *Bulk

func NewBulkFunc(tp elastictransport.Interface) NewBulk {
	_ = "STUB: not implemented"
	return *new(NewBulk)
}

func New(tp elastictransport.Interface) *Bulk { _ = "STUB: not implemented"; return nil }

func (r *Bulk) Raw(raw io.Reader) *Bulk { _ = "STUB: not implemented"; return nil }

func (r *Bulk) Request(req *Request) *Bulk { _ = "STUB: not implemented"; return nil }

func (r *Bulk) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Bulk) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Bulk) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Bulk) Header(key, value string) *Bulk { _ = "STUB: not implemented"; return nil }

func (r *Bulk) SystemId(systemid string) *Bulk { _ = "STUB: not implemented"; return nil }

func (r *Bulk) SystemApiVersion(systemapiversion string) *Bulk {
	_ = "STUB: not implemented"
	return nil
}

func (r *Bulk) Interval(duration string) *Bulk { _ = "STUB: not implemented"; return nil }

func (r *Bulk) ErrorTrace(errortrace bool) *Bulk { _ = "STUB: not implemented"; return nil }

func (r *Bulk) FilterPath(filterpaths ...string) *Bulk { _ = "STUB: not implemented"; return nil }

func (r *Bulk) Human(human bool) *Bulk { _ = "STUB: not implemented"; return nil }

func (r *Bulk) Pretty(pretty bool) *Bulk { _ = "STUB: not implemented"; return nil }
