package putazureopenai

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/azureopenaiservicetype"
)

const (
	tasktypeMask = iota + 1

	azureopenaiinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAzureopenai struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype               string
	azureopenaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutAzureopenai func(tasktype, azureopenaiinferenceid string) *PutAzureopenai

func NewPutAzureopenaiFunc(tp elastictransport.Interface) NewPutAzureopenai {
	_ = "STUB: not implemented"
	return *new(NewPutAzureopenai)
}

func New(tp elastictransport.Interface) *PutAzureopenai { _ = "STUB: not implemented"; return nil }

func (r *PutAzureopenai) Raw(raw io.Reader) *PutAzureopenai { _ = "STUB: not implemented"; return nil }

func (r *PutAzureopenai) Request(req *Request) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAzureopenai) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAzureopenai) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAzureopenai) Header(key, value string) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) _tasktype(tasktype string) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) _azureopenaiinferenceid(azureopenaiinferenceid string) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) Timeout(duration string) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) ErrorTrace(errortrace bool) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) FilterPath(filterpaths ...string) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) Human(human bool) *PutAzureopenai { _ = "STUB: not implemented"; return nil }

func (r *PutAzureopenai) Pretty(pretty bool) *PutAzureopenai { _ = "STUB: not implemented"; return nil }

func (r *PutAzureopenai) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) Service(service azureopenaiservicetype.AzureOpenAIServiceType) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) ServiceSettings(servicesettings types.AzureOpenAIServiceSettingsVariant) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureopenai) TaskSettings(tasksettings types.AzureOpenAITaskSettingsVariant) *PutAzureopenai {
	_ = "STUB: not implemented"
	return nil
}
