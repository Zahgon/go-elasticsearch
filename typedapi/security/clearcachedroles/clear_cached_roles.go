package clearcachedroles

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

type ClearCachedRoles struct {
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

type NewClearCachedRoles func(name string) *ClearCachedRoles

func NewClearCachedRolesFunc(tp elastictransport.Interface) NewClearCachedRoles {
	_ = "STUB: not implemented"
	return *new(NewClearCachedRoles)
}

func New(tp elastictransport.Interface) *ClearCachedRoles { _ = "STUB: not implemented"; return nil }

func (r *ClearCachedRoles) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedRoles) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedRoles) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedRoles) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ClearCachedRoles) Header(key, value string) *ClearCachedRoles {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRoles) _name(name string) *ClearCachedRoles {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRoles) ErrorTrace(errortrace bool) *ClearCachedRoles {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRoles) FilterPath(filterpaths ...string) *ClearCachedRoles {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRoles) Human(human bool) *ClearCachedRoles {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedRoles) Pretty(pretty bool) *ClearCachedRoles {
	_ = "STUB: not implemented"
	return nil
}
