package allocation

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catallocationcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	nodeidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Allocation struct {
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

type NewAllocation func() *Allocation

func NewAllocationFunc(tp elastictransport.Interface) NewAllocation {
	_ = "STUB: not implemented"
	return *new(NewAllocation)
}

func New(tp elastictransport.Interface) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Allocation) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Allocation) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Allocation) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Allocation) Header(key, value string) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) NodeId(nodeid string) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) H(catallocationcolumns ...catallocationcolumn.CatAllocationColumn) *Allocation {
	_ = "STUB: not implemented"
	return nil
}

func (r *Allocation) S(names ...string) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) Local(local bool) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) MasterTimeout(duration string) *Allocation {
	_ = "STUB: not implemented"
	return nil
}

func (r *Allocation) Bytes(bytes bytes.Bytes) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) Format(format string) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) Help(help bool) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) Time(time timeunit.TimeUnit) *Allocation {
	_ = "STUB: not implemented"
	return nil
}

func (r *Allocation) V(v bool) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) ErrorTrace(errortrace bool) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) FilterPath(filterpaths ...string) *Allocation {
	_ = "STUB: not implemented"
	return nil
}

func (r *Allocation) Human(human bool) *Allocation { _ = "STUB: not implemented"; return nil }

func (r *Allocation) Pretty(pretty bool) *Allocation { _ = "STUB: not implemented"; return nil }
