package getrollupindexcaps

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
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetRollupIndexCaps struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetRollupIndexCaps func(index string) *GetRollupIndexCaps

func NewGetRollupIndexCapsFunc(tp elastictransport.Interface) NewGetRollupIndexCaps {
	_ = "STUB: not implemented"
	return *new(NewGetRollupIndexCaps)
}

func New(tp elastictransport.Interface) *GetRollupIndexCaps { _ = "STUB: not implemented"; return nil }

func (r *GetRollupIndexCaps) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRollupIndexCaps) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRollupIndexCaps) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetRollupIndexCaps) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetRollupIndexCaps) Header(key, value string) *GetRollupIndexCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRollupIndexCaps) _index(index string) *GetRollupIndexCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRollupIndexCaps) ErrorTrace(errortrace bool) *GetRollupIndexCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRollupIndexCaps) FilterPath(filterpaths ...string) *GetRollupIndexCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRollupIndexCaps) Human(human bool) *GetRollupIndexCaps {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRollupIndexCaps) Pretty(pretty bool) *GetRollupIndexCaps {
	_ = "STUB: not implemented"
	return nil
}
