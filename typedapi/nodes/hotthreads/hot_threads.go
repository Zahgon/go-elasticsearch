package hotthreads

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/threadtype"
)

const (
	nodeidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type HotThreads struct {
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

type NewHotThreads func() *HotThreads

func NewHotThreadsFunc(tp elastictransport.Interface) NewHotThreads {
	_ = "STUB: not implemented"
	return *new(NewHotThreads)
}

func New(tp elastictransport.Interface) *HotThreads { _ = "STUB: not implemented"; return nil }

func (r *HotThreads) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HotThreads) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HotThreads) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r HotThreads) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *HotThreads) Header(key, value string) *HotThreads { _ = "STUB: not implemented"; return nil }

func (r *HotThreads) NodeId(nodeid string) *HotThreads { _ = "STUB: not implemented"; return nil }

func (r *HotThreads) IgnoreIdleThreads(ignoreidlethreads bool) *HotThreads {
	_ = "STUB: not implemented"
	return nil
}

func (r *HotThreads) Interval(duration string) *HotThreads { _ = "STUB: not implemented"; return nil }

func (r *HotThreads) Snapshots(snapshots string) *HotThreads { _ = "STUB: not implemented"; return nil }

func (r *HotThreads) Threads(threads string) *HotThreads { _ = "STUB: not implemented"; return nil }

func (r *HotThreads) Timeout(duration string) *HotThreads { _ = "STUB: not implemented"; return nil }

func (r *HotThreads) Type(type_ threadtype.ThreadType) *HotThreads {
	_ = "STUB: not implemented"
	return nil
}

func (r *HotThreads) Sort(sort threadtype.ThreadType) *HotThreads {
	_ = "STUB: not implemented"
	return nil
}

func (r *HotThreads) ErrorTrace(errortrace bool) *HotThreads { _ = "STUB: not implemented"; return nil }

func (r *HotThreads) FilterPath(filterpaths ...string) *HotThreads {
	_ = "STUB: not implemented"
	return nil
}

func (r *HotThreads) Human(human bool) *HotThreads { _ = "STUB: not implemented"; return nil }

func (r *HotThreads) Pretty(pretty bool) *HotThreads { _ = "STUB: not implemented"; return nil }
