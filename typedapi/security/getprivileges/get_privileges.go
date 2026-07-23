package getprivileges

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	applicationMask = iota + 1

	nameMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetPrivileges struct {
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

type NewGetPrivileges func() *GetPrivileges

func NewGetPrivilegesFunc(tp elastictransport.Interface) NewGetPrivileges {
	_ = "STUB: not implemented"
	return *new(NewGetPrivileges)
}

func New(tp elastictransport.Interface) *GetPrivileges { _ = "STUB: not implemented"; return nil }

func (r *GetPrivileges) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetPrivileges) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetPrivileges) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetPrivileges) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetPrivileges) Header(key, value string) *GetPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetPrivileges) Application(application string) *GetPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetPrivileges) Name(name string) *GetPrivileges { _ = "STUB: not implemented"; return nil }

func (r *GetPrivileges) ErrorTrace(errortrace bool) *GetPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetPrivileges) FilterPath(filterpaths ...string) *GetPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetPrivileges) Human(human bool) *GetPrivileges { _ = "STUB: not implemented"; return nil }

func (r *GetPrivileges) Pretty(pretty bool) *GetPrivileges { _ = "STUB: not implemented"; return nil }
