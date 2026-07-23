package tasks

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cattaskscolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Tasks struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewTasks func() *Tasks

func NewTasksFunc(tp elastictransport.Interface) NewTasks {
	_ = "STUB: not implemented"
	return *new(NewTasks)
}

func New(tp elastictransport.Interface) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Tasks) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Tasks) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r Tasks) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Tasks) Header(key, value string) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) Actions(actions ...string) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) Detailed(detailed bool) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) Nodes(nodes ...string) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) ParentTaskId(parenttaskid string) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) H(cattaskscolumns ...cattaskscolumn.CatTasksColumn) *Tasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *Tasks) S(names ...string) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) Timeout(duration string) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) WaitForCompletion(waitforcompletion bool) *Tasks {
	_ = "STUB: not implemented"
	return nil
}

func (r *Tasks) Bytes(bytes bytes.Bytes) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) Format(format string) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) Help(help bool) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) Time(time timeunit.TimeUnit) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) V(v bool) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) ErrorTrace(errortrace bool) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) FilterPath(filterpaths ...string) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) Human(human bool) *Tasks { _ = "STUB: not implemented"; return nil }

func (r *Tasks) Pretty(pretty bool) *Tasks { _ = "STUB: not implemented"; return nil }
