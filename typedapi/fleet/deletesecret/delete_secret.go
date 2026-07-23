package deletesecret

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

type DeleteSecret struct {
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

type NewDeleteSecret func(id string) *DeleteSecret

func NewDeleteSecretFunc(tp elastictransport.Interface) NewDeleteSecret {
	_ = "STUB: not implemented"
	return *new(NewDeleteSecret)
}

func New(tp elastictransport.Interface) *DeleteSecret { _ = "STUB: not implemented"; return nil }

func (r *DeleteSecret) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteSecret) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteSecret) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteSecret) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteSecret) Header(key, value string) *DeleteSecret {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSecret) _id(id string) *DeleteSecret { _ = "STUB: not implemented"; return nil }

func (r *DeleteSecret) ErrorTrace(errortrace bool) *DeleteSecret {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSecret) FilterPath(filterpaths ...string) *DeleteSecret {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSecret) Human(human bool) *DeleteSecret { _ = "STUB: not implemented"; return nil }

func (r *DeleteSecret) Pretty(pretty bool) *DeleteSecret { _ = "STUB: not implemented"; return nil }
