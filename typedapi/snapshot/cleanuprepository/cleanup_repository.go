package cleanuprepository

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

type CleanupRepository struct {
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

type NewCleanupRepository func(repository string) *CleanupRepository

func NewCleanupRepositoryFunc(tp elastictransport.Interface) NewCleanupRepository {
	_ = "STUB: not implemented"
	return *new(NewCleanupRepository)
}

func New(tp elastictransport.Interface) *CleanupRepository { _ = "STUB: not implemented"; return nil }

func (r *CleanupRepository) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CleanupRepository) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CleanupRepository) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CleanupRepository) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *CleanupRepository) Header(key, value string) *CleanupRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CleanupRepository) _repository(repository string) *CleanupRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CleanupRepository) MasterTimeout(duration string) *CleanupRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CleanupRepository) Timeout(duration string) *CleanupRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CleanupRepository) ErrorTrace(errortrace bool) *CleanupRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CleanupRepository) FilterPath(filterpaths ...string) *CleanupRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CleanupRepository) Human(human bool) *CleanupRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CleanupRepository) Pretty(pretty bool) *CleanupRepository {
	_ = "STUB: not implemented"
	return nil
}
