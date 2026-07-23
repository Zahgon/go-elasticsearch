package deletenode

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nodeidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteNode struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	nodeid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteNode func(nodeid string) *DeleteNode

func NewDeleteNodeFunc(tp elastictransport.Interface) NewDeleteNode {
	_ = "STUB: not implemented"
	return *new(NewDeleteNode)
}

func New(tp elastictransport.Interface) *DeleteNode { _ = "STUB: not implemented"; return nil }

func (r *DeleteNode) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteNode) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteNode) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteNode) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteNode) Header(key, value string) *DeleteNode { _ = "STUB: not implemented"; return nil }

func (r *DeleteNode) _nodeid(nodeid string) *DeleteNode { _ = "STUB: not implemented"; return nil }

func (r *DeleteNode) MasterTimeout(duration string) *DeleteNode {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteNode) Timeout(duration string) *DeleteNode { _ = "STUB: not implemented"; return nil }

func (r *DeleteNode) ErrorTrace(errortrace bool) *DeleteNode { _ = "STUB: not implemented"; return nil }

func (r *DeleteNode) FilterPath(filterpaths ...string) *DeleteNode {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteNode) Human(human bool) *DeleteNode { _ = "STUB: not implemented"; return nil }

func (r *DeleteNode) Pretty(pretty bool) *DeleteNode { _ = "STUB: not implemented"; return nil }
