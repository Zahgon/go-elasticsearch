package invalidatetoken

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

type InvalidateToken struct {
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

type NewInvalidateToken func() *InvalidateToken

func NewInvalidateTokenFunc(tp elastictransport.Interface) NewInvalidateToken {
	_ = "STUB: not implemented"
	return *new(NewInvalidateToken)
}

func New(tp elastictransport.Interface) *InvalidateToken { _ = "STUB: not implemented"; return nil }

func (r *InvalidateToken) Raw(raw io.Reader) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateToken) Request(req *Request) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateToken) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r InvalidateToken) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r InvalidateToken) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *InvalidateToken) Header(key, value string) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateToken) ErrorTrace(errortrace bool) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateToken) FilterPath(filterpaths ...string) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateToken) Human(human bool) *InvalidateToken { _ = "STUB: not implemented"; return nil }

func (r *InvalidateToken) Pretty(pretty bool) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateToken) RealmName(name string) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateToken) RefreshToken(refreshtoken string) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateToken) Token(token string) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}

func (r *InvalidateToken) Username(username string) *InvalidateToken {
	_ = "STUB: not implemented"
	return nil
}
