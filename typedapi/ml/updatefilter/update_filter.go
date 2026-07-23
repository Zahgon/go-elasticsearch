package updatefilter

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

type UpdateFilter struct {
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

type NewUpdateFilter func(filterid string) *UpdateFilter

func NewUpdateFilterFunc(tp elastictransport.Interface) NewUpdateFilter {
	_ = "STUB: not implemented"
	return *new(NewUpdateFilter)
}

func New(tp elastictransport.Interface) *UpdateFilter { _ = "STUB: not implemented"; return nil }

func (r *UpdateFilter) Raw(raw io.Reader) *UpdateFilter { _ = "STUB: not implemented"; return nil }

func (r *UpdateFilter) Request(req *Request) *UpdateFilter { _ = "STUB: not implemented"; return nil }

func (r *UpdateFilter) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateFilter) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateFilter) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateFilter) Header(key, value string) *UpdateFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilter) _filterid(filterid string) *UpdateFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilter) ErrorTrace(errortrace bool) *UpdateFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilter) FilterPath(filterpaths ...string) *UpdateFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilter) Human(human bool) *UpdateFilter { _ = "STUB: not implemented"; return nil }

func (r *UpdateFilter) Pretty(pretty bool) *UpdateFilter { _ = "STUB: not implemented"; return nil }

func (r *UpdateFilter) AddItems(additems ...string) *UpdateFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilter) Description(description string) *UpdateFilter {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilter) RemoveItems(removeitems ...string) *UpdateFilter {
	_ = "STUB: not implemented"
	return nil
}
