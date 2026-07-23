package get

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Get struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGet func(id string) *Get

func NewGetFunc(tp elastictransport.Interface) NewGet {
	_ = "STUB: not implemented"
	return *new(NewGet)
}

func New(tp elastictransport.Interface) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Get) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Get) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Get) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Get) Header(key, value string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) _id(id string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) KeepAlive(duration string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) WaitForCompletionTimeout(duration string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) ErrorTrace(errortrace bool) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) FilterPath(filterpaths ...string) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) Human(human bool) *Get { _ = "STUB: not implemented"; return nil }

func (r *Get) Pretty(pretty bool) *Get { _ = "STUB: not implemented"; return nil }
