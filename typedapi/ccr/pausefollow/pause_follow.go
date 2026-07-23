package pausefollow

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PauseFollow struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPauseFollow func(index string) *PauseFollow

func NewPauseFollowFunc(tp elastictransport.Interface) NewPauseFollow {
	_ = "STUB: not implemented"
	return *new(NewPauseFollow)
}

func New(tp elastictransport.Interface) *PauseFollow { _ = "STUB: not implemented"; return nil }

func (r *PauseFollow) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PauseFollow) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PauseFollow) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PauseFollow) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PauseFollow) Header(key, value string) *PauseFollow { _ = "STUB: not implemented"; return nil }

func (r *PauseFollow) _index(index string) *PauseFollow { _ = "STUB: not implemented"; return nil }

func (r *PauseFollow) MasterTimeout(duration string) *PauseFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseFollow) ErrorTrace(errortrace bool) *PauseFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseFollow) FilterPath(filterpaths ...string) *PauseFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *PauseFollow) Human(human bool) *PauseFollow { _ = "STUB: not implemented"; return nil }

func (r *PauseFollow) Pretty(pretty bool) *PauseFollow { _ = "STUB: not implemented"; return nil }
