package getpolicy

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

type GetPolicy struct {
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

type NewGetPolicy func() *GetPolicy

func NewGetPolicyFunc(tp elastictransport.Interface) NewGetPolicy {
	_ = "STUB: not implemented"
	return *new(NewGetPolicy)
}

func New(tp elastictransport.Interface) *GetPolicy { _ = "STUB: not implemented"; return nil }

func (r *GetPolicy) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetPolicy) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetPolicy) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetPolicy) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetPolicy) Header(key, value string) *GetPolicy { _ = "STUB: not implemented"; return nil }

func (r *GetPolicy) Name(name string) *GetPolicy { _ = "STUB: not implemented"; return nil }

func (r *GetPolicy) MasterTimeout(duration string) *GetPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetPolicy) ErrorTrace(errortrace bool) *GetPolicy { _ = "STUB: not implemented"; return nil }

func (r *GetPolicy) FilterPath(filterpaths ...string) *GetPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetPolicy) Human(human bool) *GetPolicy { _ = "STUB: not implemented"; return nil }

func (r *GetPolicy) Pretty(pretty bool) *GetPolicy { _ = "STUB: not implemented"; return nil }
