package allocationexplain

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AllocationExplain struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewAllocationExplain func() *AllocationExplain

func NewAllocationExplainFunc(tp elastictransport.Interface) NewAllocationExplain {
	_ = "STUB: not implemented"
	return *new(NewAllocationExplain)
}

func New(tp elastictransport.Interface) *AllocationExplain { _ = "STUB: not implemented"; return nil }

func (r *AllocationExplain) Raw(raw io.Reader) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) Request(req *Request) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AllocationExplain) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AllocationExplain) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *AllocationExplain) Header(key, value string) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) IncludeDiskInfo(includediskinfo bool) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) IncludeYesDecisions(includeyesdecisions bool) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) MasterTimeout(duration string) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) ErrorTrace(errortrace bool) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) FilterPath(filterpaths ...string) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) Human(human bool) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) Pretty(pretty bool) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) CurrentNode(nodeid string) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) Index(indexname string) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) Primary(primary bool) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}

func (r *AllocationExplain) Shard(shard int) *AllocationExplain {
	_ = "STUB: not implemented"
	return nil
}
