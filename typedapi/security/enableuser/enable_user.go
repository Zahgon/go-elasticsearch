package enableuser

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

type EnableUser struct {
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

type NewEnableUser func(username string) *EnableUser

func NewEnableUserFunc(tp elastictransport.Interface) NewEnableUser {
	_ = "STUB: not implemented"
	return *new(NewEnableUser)
}

func New(tp elastictransport.Interface) *EnableUser { _ = "STUB: not implemented"; return nil }

func (r *EnableUser) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnableUser) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnableUser) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnableUser) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *EnableUser) Header(key, value string) *EnableUser { _ = "STUB: not implemented"; return nil }

func (r *EnableUser) _username(username string) *EnableUser { _ = "STUB: not implemented"; return nil }

func (r *EnableUser) Refresh(refresh refresh.Refresh) *EnableUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnableUser) ErrorTrace(errortrace bool) *EnableUser { _ = "STUB: not implemented"; return nil }

func (r *EnableUser) FilterPath(filterpaths ...string) *EnableUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnableUser) Human(human bool) *EnableUser { _ = "STUB: not implemented"; return nil }

func (r *EnableUser) Pretty(pretty bool) *EnableUser { _ = "STUB: not implemented"; return nil }
