package deleteprivileges

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
	applicationMask = iota + 1

	nameMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeletePrivileges struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	application string
	name        string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeletePrivileges func(application, name string) *DeletePrivileges

func NewDeletePrivilegesFunc(tp elastictransport.Interface) NewDeletePrivileges {
	_ = "STUB: not implemented"
	return *new(NewDeletePrivileges)
}

func New(tp elastictransport.Interface) *DeletePrivileges { _ = "STUB: not implemented"; return nil }

func (r *DeletePrivileges) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeletePrivileges) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeletePrivileges) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r DeletePrivileges) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeletePrivileges) Header(key, value string) *DeletePrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePrivileges) _application(application string) *DeletePrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePrivileges) _name(name string) *DeletePrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePrivileges) Refresh(refresh refresh.Refresh) *DeletePrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePrivileges) ErrorTrace(errortrace bool) *DeletePrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePrivileges) FilterPath(filterpaths ...string) *DeletePrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePrivileges) Human(human bool) *DeletePrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePrivileges) Pretty(pretty bool) *DeletePrivileges {
	_ = "STUB: not implemented"
	return nil
}
