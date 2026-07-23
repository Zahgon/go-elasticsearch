package getrollupcaps

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetRollupCaps struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetRollupCaps func() *GetRollupCaps

func NewGetRollupCapsFunc(tp elastictransport.Interface) NewGetRollupCaps {
	_ = "STUB: not implemented"
	return *new(NewGetRollupCaps)
}

func New(tp elastictransport.Interface) *GetRollupCaps { _ = "STUB: not implemented"; return nil }

func (r *GetRollupCaps) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRollupCaps) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRollupCaps) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetRollupCaps) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetRollupCaps) Header(key, value string) *GetRollupCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRollupCaps) Id(id string) *GetRollupCaps { _ = "STUB: not implemented"; return nil }

func (r *GetRollupCaps) ErrorTrace(errortrace bool) *GetRollupCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRollupCaps) FilterPath(filterpaths ...string) *GetRollupCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRollupCaps) Human(human bool) *GetRollupCaps { _ = "STUB: not implemented"; return nil }

func (r *GetRollupCaps) Pretty(pretty bool) *GetRollupCaps { _ = "STUB: not implemented"; return nil }
