package deletealias

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1

	nameMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteAlias struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string
	name  string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteAlias func(index, name string) *DeleteAlias

func NewDeleteAliasFunc(tp elastictransport.Interface) NewDeleteAlias {
	_ = "STUB: not implemented"
	return *new(NewDeleteAlias)
}

func New(tp elastictransport.Interface) *DeleteAlias { _ = "STUB: not implemented"; return nil }

func (r *DeleteAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAlias) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAlias) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAlias) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteAlias) Header(key, value string) *DeleteAlias { _ = "STUB: not implemented"; return nil }

func (r *DeleteAlias) _index(index string) *DeleteAlias { _ = "STUB: not implemented"; return nil }

func (r *DeleteAlias) _name(name string) *DeleteAlias { _ = "STUB: not implemented"; return nil }

func (r *DeleteAlias) MasterTimeout(duration string) *DeleteAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAlias) Timeout(duration string) *DeleteAlias { _ = "STUB: not implemented"; return nil }

func (r *DeleteAlias) ErrorTrace(errortrace bool) *DeleteAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAlias) FilterPath(filterpaths ...string) *DeleteAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAlias) Human(human bool) *DeleteAlias { _ = "STUB: not implemented"; return nil }

func (r *DeleteAlias) Pretty(pretty bool) *DeleteAlias { _ = "STUB: not implemented"; return nil }
