package getuserprivileges

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetUserPrivileges struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetUserPrivileges func() *GetUserPrivileges

func NewGetUserPrivilegesFunc(tp elastictransport.Interface) NewGetUserPrivileges {
	_ = "STUB: not implemented"
	return *new(NewGetUserPrivileges)
}

func New(tp elastictransport.Interface) *GetUserPrivileges { _ = "STUB: not implemented"; return nil }

func (r *GetUserPrivileges) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetUserPrivileges) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetUserPrivileges) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetUserPrivileges) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetUserPrivileges) Header(key, value string) *GetUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUserPrivileges) ErrorTrace(errortrace bool) *GetUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUserPrivileges) FilterPath(filterpaths ...string) *GetUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUserPrivileges) Human(human bool) *GetUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetUserPrivileges) Pretty(pretty bool) *GetUserPrivileges {
	_ = "STUB: not implemented"
	return nil
}
