package authenticate

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Authenticate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewAuthenticate func() *Authenticate

func NewAuthenticateFunc(tp elastictransport.Interface) NewAuthenticate {
	_ = "STUB: not implemented"
	return *new(NewAuthenticate)
}

func New(tp elastictransport.Interface) *Authenticate { _ = "STUB: not implemented"; return nil }

func (r *Authenticate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Authenticate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Authenticate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Authenticate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Authenticate) Header(key, value string) *Authenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Authenticate) ErrorTrace(errortrace bool) *Authenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Authenticate) FilterPath(filterpaths ...string) *Authenticate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Authenticate) Human(human bool) *Authenticate { _ = "STUB: not implemented"; return nil }

func (r *Authenticate) Pretty(pretty bool) *Authenticate { _ = "STUB: not implemented"; return nil }
