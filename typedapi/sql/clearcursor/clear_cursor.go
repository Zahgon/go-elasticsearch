package clearcursor

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

type ClearCursor struct {
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

type NewClearCursor func() *ClearCursor

func NewClearCursorFunc(tp elastictransport.Interface) NewClearCursor {
	_ = "STUB: not implemented"
	return *new(NewClearCursor)
}

func New(tp elastictransport.Interface) *ClearCursor { _ = "STUB: not implemented"; return nil }

func (r *ClearCursor) Raw(raw io.Reader) *ClearCursor { _ = "STUB: not implemented"; return nil }

func (r *ClearCursor) Request(req *Request) *ClearCursor { _ = "STUB: not implemented"; return nil }

func (r *ClearCursor) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCursor) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCursor) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ClearCursor) Header(key, value string) *ClearCursor { _ = "STUB: not implemented"; return nil }

func (r *ClearCursor) ErrorTrace(errortrace bool) *ClearCursor {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCursor) FilterPath(filterpaths ...string) *ClearCursor {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCursor) Human(human bool) *ClearCursor { _ = "STUB: not implemented"; return nil }

func (r *ClearCursor) Pretty(pretty bool) *ClearCursor { _ = "STUB: not implemented"; return nil }

func (r *ClearCursor) Cursor(cursor string) *ClearCursor { _ = "STUB: not implemented"; return nil }
