package deletewatch

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

type DeleteWatch struct {
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

type NewDeleteWatch func(id string) *DeleteWatch

func NewDeleteWatchFunc(tp elastictransport.Interface) NewDeleteWatch {
	_ = "STUB: not implemented"
	return *new(NewDeleteWatch)
}

func New(tp elastictransport.Interface) *DeleteWatch { _ = "STUB: not implemented"; return nil }

func (r *DeleteWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteWatch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteWatch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteWatch) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteWatch) Header(key, value string) *DeleteWatch { _ = "STUB: not implemented"; return nil }

func (r *DeleteWatch) _id(id string) *DeleteWatch { _ = "STUB: not implemented"; return nil }

func (r *DeleteWatch) ErrorTrace(errortrace bool) *DeleteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteWatch) FilterPath(filterpaths ...string) *DeleteWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteWatch) Human(human bool) *DeleteWatch { _ = "STUB: not implemented"; return nil }

func (r *DeleteWatch) Pretty(pretty bool) *DeleteWatch { _ = "STUB: not implemented"; return nil }
