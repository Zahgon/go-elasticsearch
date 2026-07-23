package getrolemapping

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

type GetRoleMapping struct {
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

type NewGetRoleMapping func() *GetRoleMapping

func NewGetRoleMappingFunc(tp elastictransport.Interface) NewGetRoleMapping {
	_ = "STUB: not implemented"
	return *new(NewGetRoleMapping)
}

func New(tp elastictransport.Interface) *GetRoleMapping { _ = "STUB: not implemented"; return nil }

func (r *GetRoleMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRoleMapping) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRoleMapping) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetRoleMapping) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetRoleMapping) Header(key, value string) *GetRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRoleMapping) Name(name string) *GetRoleMapping { _ = "STUB: not implemented"; return nil }

func (r *GetRoleMapping) ErrorTrace(errortrace bool) *GetRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRoleMapping) FilterPath(filterpaths ...string) *GetRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRoleMapping) Human(human bool) *GetRoleMapping { _ = "STUB: not implemented"; return nil }

func (r *GetRoleMapping) Pretty(pretty bool) *GetRoleMapping { _ = "STUB: not implemented"; return nil }
