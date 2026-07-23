package deleteasync

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

type DeleteAsync struct {
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

type NewDeleteAsync func(id string) *DeleteAsync

func NewDeleteAsyncFunc(tp elastictransport.Interface) NewDeleteAsync {
	_ = "STUB: not implemented"
	return *new(NewDeleteAsync)
}

func New(tp elastictransport.Interface) *DeleteAsync { _ = "STUB: not implemented"; return nil }

func (r *DeleteAsync) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAsync) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAsync) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAsync) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteAsync) Header(key, value string) *DeleteAsync { _ = "STUB: not implemented"; return nil }

func (r *DeleteAsync) _id(id string) *DeleteAsync { _ = "STUB: not implemented"; return nil }

func (r *DeleteAsync) ErrorTrace(errortrace bool) *DeleteAsync {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAsync) FilterPath(filterpaths ...string) *DeleteAsync {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAsync) Human(human bool) *DeleteAsync { _ = "STUB: not implemented"; return nil }

func (r *DeleteAsync) Pretty(pretty bool) *DeleteAsync { _ = "STUB: not implemented"; return nil }
