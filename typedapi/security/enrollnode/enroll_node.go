package enrollnode

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type EnrollNode struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewEnrollNode func() *EnrollNode

func NewEnrollNodeFunc(tp elastictransport.Interface) NewEnrollNode {
	_ = "STUB: not implemented"
	return *new(NewEnrollNode)
}

func New(tp elastictransport.Interface) *EnrollNode { _ = "STUB: not implemented"; return nil }

func (r *EnrollNode) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnrollNode) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnrollNode) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r EnrollNode) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *EnrollNode) Header(key, value string) *EnrollNode { _ = "STUB: not implemented"; return nil }

func (r *EnrollNode) ErrorTrace(errortrace bool) *EnrollNode { _ = "STUB: not implemented"; return nil }

func (r *EnrollNode) FilterPath(filterpaths ...string) *EnrollNode {
	_ = "STUB: not implemented"
	return nil
}

func (r *EnrollNode) Human(human bool) *EnrollNode { _ = "STUB: not implemented"; return nil }

func (r *EnrollNode) Pretty(pretty bool) *EnrollNode { _ = "STUB: not implemented"; return nil }
