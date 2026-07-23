package threadpool

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catthreadpoolcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	threadpoolpatternsMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ThreadPool struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	threadpoolpatterns string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewThreadPool func() *ThreadPool

func NewThreadPoolFunc(tp elastictransport.Interface) NewThreadPool {
	_ = "STUB: not implemented"
	return *new(NewThreadPool)
}

func New(tp elastictransport.Interface) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ThreadPool) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ThreadPool) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r ThreadPool) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ThreadPool) Header(key, value string) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) ThreadPoolPatterns(threadpoolpatterns string) *ThreadPool {
	_ = "STUB: not implemented"
	return nil
}

func (r *ThreadPool) H(catthreadpoolcolumns ...catthreadpoolcolumn.CatThreadPoolColumn) *ThreadPool {
	_ = "STUB: not implemented"
	return nil
}

func (r *ThreadPool) S(names ...string) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) Local(local bool) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) MasterTimeout(duration string) *ThreadPool {
	_ = "STUB: not implemented"
	return nil
}

func (r *ThreadPool) Bytes(bytes bytes.Bytes) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) Format(format string) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) Help(help bool) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) Time(time timeunit.TimeUnit) *ThreadPool {
	_ = "STUB: not implemented"
	return nil
}

func (r *ThreadPool) V(v bool) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) ErrorTrace(errortrace bool) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) FilterPath(filterpaths ...string) *ThreadPool {
	_ = "STUB: not implemented"
	return nil
}

func (r *ThreadPool) Human(human bool) *ThreadPool { _ = "STUB: not implemented"; return nil }

func (r *ThreadPool) Pretty(pretty bool) *ThreadPool { _ = "STUB: not implemented"; return nil }
