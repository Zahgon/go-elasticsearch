package secretget

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

type SecretGet struct {
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

type NewSecretGet func(id string) *SecretGet

func NewSecretGetFunc(tp elastictransport.Interface) NewSecretGet {
	_ = "STUB: not implemented"
	return *new(NewSecretGet)
}

func New(tp elastictransport.Interface) *SecretGet { _ = "STUB: not implemented"; return nil }

func (r *SecretGet) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretGet) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretGet) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SecretGet) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SecretGet) Header(key, value string) *SecretGet { _ = "STUB: not implemented"; return nil }

func (r *SecretGet) _id(id string) *SecretGet { _ = "STUB: not implemented"; return nil }

func (r *SecretGet) ErrorTrace(errortrace bool) *SecretGet { _ = "STUB: not implemented"; return nil }

func (r *SecretGet) FilterPath(filterpaths ...string) *SecretGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *SecretGet) Human(human bool) *SecretGet { _ = "STUB: not implemented"; return nil }

func (r *SecretGet) Pretty(pretty bool) *SecretGet { _ = "STUB: not implemented"; return nil }
