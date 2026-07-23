package putllama

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/llamaservicetype"
)

const (
	tasktypeMask = iota + 1

	llamainferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutLlama struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype         string
	llamainferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutLlama func(tasktype, llamainferenceid string) *PutLlama

func NewPutLlamaFunc(tp elastictransport.Interface) NewPutLlama {
	_ = "STUB: not implemented"
	return *new(NewPutLlama)
}

func New(tp elastictransport.Interface) *PutLlama { _ = "STUB: not implemented"; return nil }

func (r *PutLlama) Raw(raw io.Reader) *PutLlama { _ = "STUB: not implemented"; return nil }

func (r *PutLlama) Request(req *Request) *PutLlama { _ = "STUB: not implemented"; return nil }

func (r *PutLlama) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutLlama) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutLlama) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutLlama) Header(key, value string) *PutLlama { _ = "STUB: not implemented"; return nil }

func (r *PutLlama) _tasktype(tasktype string) *PutLlama { _ = "STUB: not implemented"; return nil }

func (r *PutLlama) _llamainferenceid(llamainferenceid string) *PutLlama {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLlama) Timeout(duration string) *PutLlama { _ = "STUB: not implemented"; return nil }

func (r *PutLlama) ErrorTrace(errortrace bool) *PutLlama { _ = "STUB: not implemented"; return nil }

func (r *PutLlama) FilterPath(filterpaths ...string) *PutLlama {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLlama) Human(human bool) *PutLlama { _ = "STUB: not implemented"; return nil }

func (r *PutLlama) Pretty(pretty bool) *PutLlama { _ = "STUB: not implemented"; return nil }

func (r *PutLlama) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutLlama {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLlama) Service(service llamaservicetype.LlamaServiceType) *PutLlama {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutLlama) ServiceSettings(servicesettings types.LlamaServiceSettingsVariant) *PutLlama {
	_ = "STUB: not implemented"
	return nil
}
