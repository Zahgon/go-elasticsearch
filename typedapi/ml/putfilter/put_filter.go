package putfilter

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	filteridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutFilter struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	filterid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutFilter func(filterid string) *PutFilter

func NewPutFilterFunc(tp elastictransport.Interface) NewPutFilter {
	_ = "STUB: not implemented"
	return *new(NewPutFilter)
}

func New(tp elastictransport.Interface) *PutFilter { _ = "STUB: not implemented"; return nil }

func (r *PutFilter) Raw(raw io.Reader) *PutFilter { _ = "STUB: not implemented"; return nil }

func (r *PutFilter) Request(req *Request) *PutFilter { _ = "STUB: not implemented"; return nil }

func (r *PutFilter) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutFilter) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutFilter) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutFilter) Header(key, value string) *PutFilter { _ = "STUB: not implemented"; return nil }

func (r *PutFilter) _filterid(filterid string) *PutFilter { _ = "STUB: not implemented"; return nil }

func (r *PutFilter) ErrorTrace(errortrace bool) *PutFilter { _ = "STUB: not implemented"; return nil }

func (r *PutFilter) FilterPath(filterpaths ...string) *PutFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFilter) Human(human bool) *PutFilter { _ = "STUB: not implemented"; return nil }

func (r *PutFilter) Pretty(pretty bool) *PutFilter { _ = "STUB: not implemented"; return nil }

func (r *PutFilter) Description(description string) *PutFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFilter) Items(items ...string) *PutFilter { _ = "STUB: not implemented"; return nil }
