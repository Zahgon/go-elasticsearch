package deleterole

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

type DeleteRole struct {
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

type NewDeleteRole func(name string) *DeleteRole

func NewDeleteRoleFunc(tp elastictransport.Interface) NewDeleteRole {
	_ = "STUB: not implemented"
	return *new(NewDeleteRole)
}

func New(tp elastictransport.Interface) *DeleteRole { _ = "STUB: not implemented"; return nil }

func (r *DeleteRole) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRole) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRole) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRole) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteRole) Header(key, value string) *DeleteRole { _ = "STUB: not implemented"; return nil }

func (r *DeleteRole) _name(name string) *DeleteRole { _ = "STUB: not implemented"; return nil }

func (r *DeleteRole) Refresh(refresh refresh.Refresh) *DeleteRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRole) ErrorTrace(errortrace bool) *DeleteRole { _ = "STUB: not implemented"; return nil }

func (r *DeleteRole) FilterPath(filterpaths ...string) *DeleteRole {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRole) Human(human bool) *DeleteRole { _ = "STUB: not implemented"; return nil }

func (r *DeleteRole) Pretty(pretty bool) *DeleteRole { _ = "STUB: not implemented"; return nil }
