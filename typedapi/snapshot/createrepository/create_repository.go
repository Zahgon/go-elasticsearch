package createrepository

import (
	gobytes "bytes"
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

type CreateRepository struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	repository string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCreateRepository func(repository string) *CreateRepository

func NewCreateRepositoryFunc(tp elastictransport.Interface) NewCreateRepository {
	_ = "STUB: not implemented"
	return *new(NewCreateRepository)
}

func New(tp elastictransport.Interface) *CreateRepository { _ = "STUB: not implemented"; return nil }

func (r *CreateRepository) Raw(raw io.Reader) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) Request(req *Request) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateRepository) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CreateRepository) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CreateRepository) Header(key, value string) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) _repository(repository string) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) MasterTimeout(duration string) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) Timeout(duration string) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) Verify(verify bool) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) ErrorTrace(errortrace bool) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) FilterPath(filterpaths ...string) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) Human(human bool) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}

func (r *CreateRepository) Pretty(pretty bool) *CreateRepository {
	_ = "STUB: not implemented"
	return nil
}
