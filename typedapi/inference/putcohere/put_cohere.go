package putcohere

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cohereservicetype"
)

const (
	tasktypeMask = iota + 1

	cohereinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutCohere struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype          string
	cohereinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutCohere func(tasktype, cohereinferenceid string) *PutCohere

func NewPutCohereFunc(tp elastictransport.Interface) NewPutCohere {
	_ = "STUB: not implemented"
	return *new(NewPutCohere)
}

func New(tp elastictransport.Interface) *PutCohere { _ = "STUB: not implemented"; return nil }

func (r *PutCohere) Raw(raw io.Reader) *PutCohere { _ = "STUB: not implemented"; return nil }

func (r *PutCohere) Request(req *Request) *PutCohere { _ = "STUB: not implemented"; return nil }

func (r *PutCohere) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutCohere) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutCohere) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutCohere) Header(key, value string) *PutCohere { _ = "STUB: not implemented"; return nil }

func (r *PutCohere) _tasktype(tasktype string) *PutCohere { _ = "STUB: not implemented"; return nil }

func (r *PutCohere) _cohereinferenceid(cohereinferenceid string) *PutCohere {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCohere) Timeout(duration string) *PutCohere { _ = "STUB: not implemented"; return nil }

func (r *PutCohere) ErrorTrace(errortrace bool) *PutCohere { _ = "STUB: not implemented"; return nil }

func (r *PutCohere) FilterPath(filterpaths ...string) *PutCohere {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCohere) Human(human bool) *PutCohere { _ = "STUB: not implemented"; return nil }

func (r *PutCohere) Pretty(pretty bool) *PutCohere { _ = "STUB: not implemented"; return nil }

func (r *PutCohere) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutCohere {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCohere) Service(service cohereservicetype.CohereServiceType) *PutCohere {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCohere) ServiceSettings(servicesettings types.CohereServiceSettingsVariant) *PutCohere {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCohere) TaskSettings(tasksettings types.CohereTaskSettingsVariant) *PutCohere {
	_ = "STUB: not implemented"
	return nil
}
