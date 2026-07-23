package disableuser

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
	usernameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DisableUser struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	username string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDisableUser func(username string) *DisableUser

func NewDisableUserFunc(tp elastictransport.Interface) NewDisableUser {
	_ = "STUB: not implemented"
	return *new(NewDisableUser)
}

func New(tp elastictransport.Interface) *DisableUser { _ = "STUB: not implemented"; return nil }

func (r *DisableUser) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DisableUser) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DisableUser) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DisableUser) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DisableUser) Header(key, value string) *DisableUser { _ = "STUB: not implemented"; return nil }

func (r *DisableUser) _username(username string) *DisableUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUser) Refresh(refresh refresh.Refresh) *DisableUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUser) ErrorTrace(errortrace bool) *DisableUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUser) FilterPath(filterpaths ...string) *DisableUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *DisableUser) Human(human bool) *DisableUser { _ = "STUB: not implemented"; return nil }

func (r *DisableUser) Pretty(pretty bool) *DisableUser { _ = "STUB: not implemented"; return nil }
