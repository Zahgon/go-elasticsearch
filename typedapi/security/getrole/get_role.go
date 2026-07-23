package getrole

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetRole struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetRole func() *GetRole

func NewGetRoleFunc(tp elastictransport.Interface) NewGetRole {
	_ = "STUB: not implemented"
	return *new(NewGetRole)
}

func New(tp elastictransport.Interface) *GetRole { _ = "STUB: not implemented"; return nil }

func (r *GetRole) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRole) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRole) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetRole) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetRole) Header(key, value string) *GetRole { _ = "STUB: not implemented"; return nil }

func (r *GetRole) Name(name string) *GetRole { _ = "STUB: not implemented"; return nil }

func (r *GetRole) IncludeImplicit(includeimplicit bool) *GetRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRole) ErrorTrace(errortrace bool) *GetRole { _ = "STUB: not implemented"; return nil }

func (r *GetRole) FilterPath(filterpaths ...string) *GetRole { _ = "STUB: not implemented"; return nil }

func (r *GetRole) Human(human bool) *GetRole { _ = "STUB: not implemented"; return nil }

func (r *GetRole) Pretty(pretty bool) *GetRole { _ = "STUB: not implemented"; return nil }
