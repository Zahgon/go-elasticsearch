package verifyrepository

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

type VerifyRepository struct {
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

type NewVerifyRepository func(repository string) *VerifyRepository

func NewVerifyRepositoryFunc(tp elastictransport.Interface) NewVerifyRepository {
	_ = "STUB: not implemented"
	return *new(NewVerifyRepository)
}

func New(tp elastictransport.Interface) *VerifyRepository { _ = "STUB: not implemented"; return nil }

func (r *VerifyRepository) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r VerifyRepository) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r VerifyRepository) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r VerifyRepository) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *VerifyRepository) Header(key, value string) *VerifyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *VerifyRepository) _repository(repository string) *VerifyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *VerifyRepository) MasterTimeout(duration string) *VerifyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *VerifyRepository) Timeout(duration string) *VerifyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *VerifyRepository) ErrorTrace(errortrace bool) *VerifyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *VerifyRepository) FilterPath(filterpaths ...string) *VerifyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *VerifyRepository) Human(human bool) *VerifyRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *VerifyRepository) Pretty(pretty bool) *VerifyRepository {
	_ = "STUB: not implemented"
	return nil
}
