package putview

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
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutView struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutView func(name string) *PutView

func NewPutViewFunc(tp elastictransport.Interface) NewPutView {
	_ = "STUB: not implemented"
	return *new(NewPutView)
}

func New(tp elastictransport.Interface) *PutView { _ = "STUB: not implemented"; return nil }

func (r *PutView) Raw(raw io.Reader) *PutView { _ = "STUB: not implemented"; return nil }

func (r *PutView) Request(req *Request) *PutView { _ = "STUB: not implemented"; return nil }

func (r *PutView) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutView) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutView) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutView) Header(key, value string) *PutView { _ = "STUB: not implemented"; return nil }

func (r *PutView) _name(name string) *PutView { _ = "STUB: not implemented"; return nil }

func (r *PutView) ErrorTrace(errortrace bool) *PutView { _ = "STUB: not implemented"; return nil }

func (r *PutView) FilterPath(filterpaths ...string) *PutView { _ = "STUB: not implemented"; return nil }

func (r *PutView) Human(human bool) *PutView { _ = "STUB: not implemented"; return nil }

func (r *PutView) Pretty(pretty bool) *PutView { _ = "STUB: not implemented"; return nil }

func (r *PutView) Query(query string) *PutView { _ = "STUB: not implemented"; return nil }
