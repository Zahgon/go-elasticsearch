package putnvidia

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/nvidiaservicetype"
)

const (
	tasktypeMask = iota + 1

	nvidiainferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutNvidia struct {
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
	nvidiainferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutNvidia func(tasktype, nvidiainferenceid string) *PutNvidia

func NewPutNvidiaFunc(tp elastictransport.Interface) NewPutNvidia {
	_ = "STUB: not implemented"
	return *new(NewPutNvidia)
}

func New(tp elastictransport.Interface) *PutNvidia { _ = "STUB: not implemented"; return nil }

func (r *PutNvidia) Raw(raw io.Reader) *PutNvidia { _ = "STUB: not implemented"; return nil }

func (r *PutNvidia) Request(req *Request) *PutNvidia { _ = "STUB: not implemented"; return nil }

func (r *PutNvidia) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutNvidia) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutNvidia) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutNvidia) Header(key, value string) *PutNvidia { _ = "STUB: not implemented"; return nil }

func (r *PutNvidia) _tasktype(tasktype string) *PutNvidia { _ = "STUB: not implemented"; return nil }

func (r *PutNvidia) _nvidiainferenceid(nvidiainferenceid string) *PutNvidia {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutNvidia) Timeout(duration string) *PutNvidia { _ = "STUB: not implemented"; return nil }

func (r *PutNvidia) ErrorTrace(errortrace bool) *PutNvidia { _ = "STUB: not implemented"; return nil }

func (r *PutNvidia) FilterPath(filterpaths ...string) *PutNvidia {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutNvidia) Human(human bool) *PutNvidia { _ = "STUB: not implemented"; return nil }

func (r *PutNvidia) Pretty(pretty bool) *PutNvidia { _ = "STUB: not implemented"; return nil }

func (r *PutNvidia) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutNvidia {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutNvidia) Service(service nvidiaservicetype.NvidiaServiceType) *PutNvidia {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutNvidia) ServiceSettings(servicesettings types.NvidiaServiceSettingsVariant) *PutNvidia {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutNvidia) TaskSettings(tasksettings types.NvidiaTaskSettingsVariant) *PutNvidia {
	_ = "STUB: not implemented"
	return nil
}
