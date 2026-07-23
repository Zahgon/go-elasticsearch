package getbuiltinprivileges

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetBuiltinPrivileges struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetBuiltinPrivileges func() *GetBuiltinPrivileges

func NewGetBuiltinPrivilegesFunc(tp elastictransport.Interface) NewGetBuiltinPrivileges {
	_ = "STUB: not implemented"
	return *new(NewGetBuiltinPrivileges)
}

func New(tp elastictransport.Interface) *GetBuiltinPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuiltinPrivileges) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBuiltinPrivileges) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBuiltinPrivileges) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBuiltinPrivileges) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetBuiltinPrivileges) Header(key, value string) *GetBuiltinPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuiltinPrivileges) ErrorTrace(errortrace bool) *GetBuiltinPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuiltinPrivileges) FilterPath(filterpaths ...string) *GetBuiltinPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuiltinPrivileges) Human(human bool) *GetBuiltinPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuiltinPrivileges) Pretty(pretty bool) *GetBuiltinPrivileges {
	_ = "STUB: not implemented"
	return nil
}
