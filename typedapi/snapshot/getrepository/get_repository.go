package getrepository

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

type GetRepository struct {
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

type NewGetRepository func() *GetRepository

func NewGetRepositoryFunc(tp elastictransport.Interface) NewGetRepository {
	_ = "STUB: not implemented"
	return *new(NewGetRepository)
}

func New(tp elastictransport.Interface) *GetRepository { _ = "STUB: not implemented"; return nil }

func (r *GetRepository) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRepository) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRepository) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetRepository) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetRepository) Header(key, value string) *GetRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepository) Repository(repository string) *GetRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepository) Local(local bool) *GetRepository { _ = "STUB: not implemented"; return nil }

func (r *GetRepository) MasterTimeout(duration string) *GetRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepository) ErrorTrace(errortrace bool) *GetRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepository) FilterPath(filterpaths ...string) *GetRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepository) Human(human bool) *GetRepository { _ = "STUB: not implemented"; return nil }

func (r *GetRepository) Pretty(pretty bool) *GetRepository { _ = "STUB: not implemented"; return nil }
