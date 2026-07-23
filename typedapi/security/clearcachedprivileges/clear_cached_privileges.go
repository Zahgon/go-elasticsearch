package clearcachedprivileges

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
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ClearCachedPrivileges struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	application string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewClearCachedPrivileges func(application string) *ClearCachedPrivileges

func NewClearCachedPrivilegesFunc(tp elastictransport.Interface) NewClearCachedPrivileges {
	_ = "STUB: not implemented"
	return *new(NewClearCachedPrivileges)
}

func New(tp elastictransport.Interface) *ClearCachedPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedPrivileges) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedPrivileges) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedPrivileges) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCachedPrivileges) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ClearCachedPrivileges) Header(key, value string) *ClearCachedPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedPrivileges) _application(application string) *ClearCachedPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedPrivileges) ErrorTrace(errortrace bool) *ClearCachedPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedPrivileges) FilterPath(filterpaths ...string) *ClearCachedPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedPrivileges) Human(human bool) *ClearCachedPrivileges {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCachedPrivileges) Pretty(pretty bool) *ClearCachedPrivileges {
	_ = "STUB: not implemented"
	return nil
}
