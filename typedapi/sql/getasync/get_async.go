package getasync

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

type GetAsync struct {
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

type NewGetAsync func(id string) *GetAsync

func NewGetAsyncFunc(tp elastictransport.Interface) NewGetAsync {
	_ = "STUB: not implemented"
	return *new(NewGetAsync)
}

func New(tp elastictransport.Interface) *GetAsync { _ = "STUB: not implemented"; return nil }

func (r *GetAsync) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAsync) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAsync) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAsync) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetAsync) Header(key, value string) *GetAsync { _ = "STUB: not implemented"; return nil }

func (r *GetAsync) _id(id string) *GetAsync { _ = "STUB: not implemented"; return nil }

func (r *GetAsync) Delimiter(delimiter string) *GetAsync { _ = "STUB: not implemented"; return nil }

func (r *GetAsync) Format(format string) *GetAsync { _ = "STUB: not implemented"; return nil }

func (r *GetAsync) KeepAlive(duration string) *GetAsync { _ = "STUB: not implemented"; return nil }

func (r *GetAsync) WaitForCompletionTimeout(duration string) *GetAsync {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAsync) ErrorTrace(errortrace bool) *GetAsync { _ = "STUB: not implemented"; return nil }

func (r *GetAsync) FilterPath(filterpaths ...string) *GetAsync {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAsync) Human(human bool) *GetAsync { _ = "STUB: not implemented"; return nil }

func (r *GetAsync) Pretty(pretty bool) *GetAsync { _ = "STUB: not implemented"; return nil }
