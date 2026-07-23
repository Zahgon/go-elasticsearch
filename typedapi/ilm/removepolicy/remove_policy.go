package removepolicy

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

type RemovePolicy struct {
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

type NewRemovePolicy func(index string) *RemovePolicy

func NewRemovePolicyFunc(tp elastictransport.Interface) NewRemovePolicy {
	_ = "STUB: not implemented"
	return *new(NewRemovePolicy)
}

func New(tp elastictransport.Interface) *RemovePolicy { _ = "STUB: not implemented"; return nil }

func (r *RemovePolicy) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RemovePolicy) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RemovePolicy) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RemovePolicy) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *RemovePolicy) Header(key, value string) *RemovePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemovePolicy) _index(index string) *RemovePolicy { _ = "STUB: not implemented"; return nil }

func (r *RemovePolicy) ErrorTrace(errortrace bool) *RemovePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemovePolicy) FilterPath(filterpaths ...string) *RemovePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemovePolicy) Human(human bool) *RemovePolicy { _ = "STUB: not implemented"; return nil }

func (r *RemovePolicy) Pretty(pretty bool) *RemovePolicy { _ = "STUB: not implemented"; return nil }
