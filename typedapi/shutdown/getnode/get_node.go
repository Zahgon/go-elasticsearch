package getnode

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

type GetNode struct {
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

type NewGetNode func() *GetNode

func NewGetNodeFunc(tp elastictransport.Interface) NewGetNode {
	_ = "STUB: not implemented"
	return *new(NewGetNode)
}

func New(tp elastictransport.Interface) *GetNode { _ = "STUB: not implemented"; return nil }

func (r *GetNode) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetNode) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetNode) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetNode) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetNode) Header(key, value string) *GetNode { _ = "STUB: not implemented"; return nil }

func (r *GetNode) NodeId(nodeid string) *GetNode { _ = "STUB: not implemented"; return nil }

func (r *GetNode) MasterTimeout(duration string) *GetNode { _ = "STUB: not implemented"; return nil }

func (r *GetNode) ErrorTrace(errortrace bool) *GetNode { _ = "STUB: not implemented"; return nil }

func (r *GetNode) FilterPath(filterpaths ...string) *GetNode { _ = "STUB: not implemented"; return nil }

func (r *GetNode) Human(human bool) *GetNode { _ = "STUB: not implemented"; return nil }

func (r *GetNode) Pretty(pretty bool) *GetNode { _ = "STUB: not implemented"; return nil }
