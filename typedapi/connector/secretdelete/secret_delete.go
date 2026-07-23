package secretdelete

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

type SecretDelete struct {
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

type NewSecretDelete func(id string) *SecretDelete

func NewSecretDeleteFunc(tp elastictransport.Interface) NewSecretDelete {
	_ = "STUB: not implemented"
	return *new(NewSecretDelete)
}

func New(tp elastictransport.Interface) *SecretDelete { _ = "STUB: not implemented"; return nil }

func (r *SecretDelete) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretDelete) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretDelete) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretDelete) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SecretDelete) Header(key, value string) *SecretDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *SecretDelete) _id(id string) *SecretDelete { _ = "STUB: not implemented"; return nil }

func (r *SecretDelete) ErrorTrace(errortrace bool) *SecretDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *SecretDelete) FilterPath(filterpaths ...string) *SecretDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *SecretDelete) Human(human bool) *SecretDelete { _ = "STUB: not implemented"; return nil }

func (r *SecretDelete) Pretty(pretty bool) *SecretDelete { _ = "STUB: not implemented"; return nil }
