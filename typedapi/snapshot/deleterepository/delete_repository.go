package deleterepository

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	repositoryMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteRepository struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	repository string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteRepository func(repository string) *DeleteRepository

func NewDeleteRepositoryFunc(tp elastictransport.Interface) NewDeleteRepository {
	_ = "STUB: not implemented"
	return *new(NewDeleteRepository)
}

func New(tp elastictransport.Interface) *DeleteRepository { _ = "STUB: not implemented"; return nil }

func (r *DeleteRepository) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRepository) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRepository) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRepository) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteRepository) Header(key, value string) *DeleteRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRepository) _repository(repository string) *DeleteRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRepository) MasterTimeout(duration string) *DeleteRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRepository) Timeout(duration string) *DeleteRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRepository) ErrorTrace(errortrace bool) *DeleteRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRepository) FilterPath(filterpaths ...string) *DeleteRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRepository) Human(human bool) *DeleteRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRepository) Pretty(pretty bool) *DeleteRepository {
	_ = "STUB: not implemented"
	return nil
}
