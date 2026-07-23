package deleterolemapping

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteRoleMapping struct {
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

type NewDeleteRoleMapping func(name string) *DeleteRoleMapping

func NewDeleteRoleMappingFunc(tp elastictransport.Interface) NewDeleteRoleMapping {
	_ = "STUB: not implemented"
	return *new(NewDeleteRoleMapping)
}

func New(tp elastictransport.Interface) *DeleteRoleMapping { _ = "STUB: not implemented"; return nil }

func (r *DeleteRoleMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRoleMapping) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRoleMapping) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRoleMapping) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteRoleMapping) Header(key, value string) *DeleteRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRoleMapping) _name(name string) *DeleteRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRoleMapping) Refresh(refresh refresh.Refresh) *DeleteRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRoleMapping) ErrorTrace(errortrace bool) *DeleteRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRoleMapping) FilterPath(filterpaths ...string) *DeleteRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRoleMapping) Human(human bool) *DeleteRoleMapping {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRoleMapping) Pretty(pretty bool) *DeleteRoleMapping {
	_ = "STUB: not implemented"
	return nil
}
