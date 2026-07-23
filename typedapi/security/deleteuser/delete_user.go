package deleteuser

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

type DeleteUser struct {
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

type NewDeleteUser func(username string) *DeleteUser

func NewDeleteUserFunc(tp elastictransport.Interface) NewDeleteUser {
	_ = "STUB: not implemented"
	return *new(NewDeleteUser)
}

func New(tp elastictransport.Interface) *DeleteUser { _ = "STUB: not implemented"; return nil }

func (r *DeleteUser) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteUser) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteUser) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteUser) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteUser) Header(key, value string) *DeleteUser { _ = "STUB: not implemented"; return nil }

func (r *DeleteUser) _username(username string) *DeleteUser { _ = "STUB: not implemented"; return nil }

func (r *DeleteUser) Refresh(refresh refresh.Refresh) *DeleteUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteUser) ErrorTrace(errortrace bool) *DeleteUser { _ = "STUB: not implemented"; return nil }

func (r *DeleteUser) FilterPath(filterpaths ...string) *DeleteUser {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteUser) Human(human bool) *DeleteUser { _ = "STUB: not implemented"; return nil }

func (r *DeleteUser) Pretty(pretty bool) *DeleteUser { _ = "STUB: not implemented"; return nil }
