package ping

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Ping struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPing func() *Ping

func NewPingFunc(tp elastictransport.Interface) NewPing {
	_ = "STUB: not implemented"
	return *new(NewPing)
}

func New(tp elastictransport.Interface) *Ping { _ = "STUB: not implemented"; return nil }

func (r *Ping) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Ping) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Ping) Do(ctx context.Context) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r Ping) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Ping) Header(key, value string) *Ping { _ = "STUB: not implemented"; return nil }

func (r *Ping) ErrorTrace(errortrace bool) *Ping { _ = "STUB: not implemented"; return nil }

func (r *Ping) FilterPath(filterpaths ...string) *Ping { _ = "STUB: not implemented"; return nil }

func (r *Ping) Human(human bool) *Ping { _ = "STUB: not implemented"; return nil }

func (r *Ping) Pretty(pretty bool) *Ping { _ = "STUB: not implemented"; return nil }
