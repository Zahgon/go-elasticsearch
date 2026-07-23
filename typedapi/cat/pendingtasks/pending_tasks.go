package pendingtasks

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catpendingtaskscolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PendingTasks struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPendingTasks func() *PendingTasks

func NewPendingTasksFunc(tp elastictransport.Interface) NewPendingTasks {
	_ = "STUB: not implemented"
	return *new(NewPendingTasks)
}

func New(tp elastictransport.Interface) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PendingTasks) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PendingTasks) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r PendingTasks) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PendingTasks) Header(key, value string) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) H(catpendingtaskscolumns ...catpendingtaskscolumn.CatPendingTasksColumn) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) S(names ...string) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) Local(local bool) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) MasterTimeout(duration string) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) Bytes(bytes bytes.Bytes) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) Format(format string) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) Help(help bool) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) Time(time timeunit.TimeUnit) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) V(v bool) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) ErrorTrace(errortrace bool) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) FilterPath(filterpaths ...string) *PendingTasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *PendingTasks) Human(human bool) *PendingTasks { _ = "STUB: not implemented"; return nil }

func (r *PendingTasks) Pretty(pretty bool) *PendingTasks { _ = "STUB: not implemented"; return nil }
