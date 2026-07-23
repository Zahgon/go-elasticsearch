package getlifecycle

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	policyidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	policyid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetLifecycle func() *GetLifecycle

func NewGetLifecycleFunc(tp elastictransport.Interface) NewGetLifecycle {
	_ = "STUB: not implemented"
	return *new(NewGetLifecycle)
}

func New(tp elastictransport.Interface) *GetLifecycle { _ = "STUB: not implemented"; return nil }

func (r *GetLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetLifecycle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetLifecycle) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetLifecycle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetLifecycle) Header(key, value string) *GetLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetLifecycle) PolicyId(policyid string) *GetLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetLifecycle) MasterTimeout(duration string) *GetLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetLifecycle) Timeout(duration string) *GetLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetLifecycle) ErrorTrace(errortrace bool) *GetLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetLifecycle) FilterPath(filterpaths ...string) *GetLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetLifecycle) Human(human bool) *GetLifecycle { _ = "STUB: not implemented"; return nil }

func (r *GetLifecycle) Pretty(pretty bool) *GetLifecycle { _ = "STUB: not implemented"; return nil }
