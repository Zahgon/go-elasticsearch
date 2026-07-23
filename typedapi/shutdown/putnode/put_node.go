package putnode

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/type_"
)

const (
	nodeidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutNode struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	nodeid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutNode func(nodeid string) *PutNode

func NewPutNodeFunc(tp elastictransport.Interface) NewPutNode {
	_ = "STUB: not implemented"
	return *new(NewPutNode)
}

func New(tp elastictransport.Interface) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) Raw(raw io.Reader) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) Request(req *Request) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutNode) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutNode) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutNode) Header(key, value string) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) _nodeid(nodeid string) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) MasterTimeout(duration string) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) Timeout(duration string) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) ErrorTrace(errortrace bool) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) FilterPath(filterpaths ...string) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) Human(human bool) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) Pretty(pretty bool) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) AllocationDelay(allocationdelay string) *PutNode {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutNode) Reason(reason string) *PutNode { _ = "STUB: not implemented"; return nil }

func (r *PutNode) TargetNodeName(targetnodename string) *PutNode {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutNode) Type(type_ type_.Type) *PutNode { _ = "STUB: not implemented"; return nil }
