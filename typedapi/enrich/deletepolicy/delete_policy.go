package deletepolicy

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeletePolicy struct {
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

type NewDeletePolicy func(name string) *DeletePolicy

func NewDeletePolicyFunc(tp elastictransport.Interface) NewDeletePolicy {
	_ = "STUB: not implemented"
	return *new(NewDeletePolicy)
}

func New(tp elastictransport.Interface) *DeletePolicy { _ = "STUB: not implemented"; return nil }

func (r *DeletePolicy) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeletePolicy) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeletePolicy) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeletePolicy) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeletePolicy) Header(key, value string) *DeletePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePolicy) _name(name string) *DeletePolicy { _ = "STUB: not implemented"; return nil }

func (r *DeletePolicy) MasterTimeout(duration string) *DeletePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePolicy) ErrorTrace(errortrace bool) *DeletePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePolicy) FilterPath(filterpaths ...string) *DeletePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePolicy) Human(human bool) *DeletePolicy { _ = "STUB: not implemented"; return nil }

func (r *DeletePolicy) Pretty(pretty bool) *DeletePolicy { _ = "STUB: not implemented"; return nil }
