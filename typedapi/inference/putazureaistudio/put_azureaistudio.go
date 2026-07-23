package putazureaistudio

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/azureaistudioservicetype"
)

const (
	tasktypeMask = iota + 1

	azureaistudioinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAzureaistudio struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                 string
	azureaistudioinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutAzureaistudio func(tasktype, azureaistudioinferenceid string) *PutAzureaistudio

func NewPutAzureaistudioFunc(tp elastictransport.Interface) NewPutAzureaistudio {
	_ = "STUB: not implemented"
	return *new(NewPutAzureaistudio)
}

func New(tp elastictransport.Interface) *PutAzureaistudio { _ = "STUB: not implemented"; return nil }

func (r *PutAzureaistudio) Raw(raw io.Reader) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) Request(req *Request) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAzureaistudio) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAzureaistudio) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAzureaistudio) Header(key, value string) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) _tasktype(tasktype string) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) _azureaistudioinferenceid(azureaistudioinferenceid string) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) Timeout(duration string) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) ErrorTrace(errortrace bool) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) FilterPath(filterpaths ...string) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) Human(human bool) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) Pretty(pretty bool) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) Service(service azureaistudioservicetype.AzureAiStudioServiceType) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) ServiceSettings(servicesettings types.AzureAiStudioServiceSettingsVariant) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAzureaistudio) TaskSettings(tasksettings types.AzureAiStudioTaskSettingsVariant) *PutAzureaistudio {
	_ = "STUB: not implemented"
	return nil
}
