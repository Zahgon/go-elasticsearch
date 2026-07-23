package putopenai

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openaiservicetype"
)

const (
	tasktypeMask = iota + 1

	openaiinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutOpenai struct {
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
	openaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutOpenai func(tasktype, openaiinferenceid string) *PutOpenai

func NewPutOpenaiFunc(tp elastictransport.Interface) NewPutOpenai {
	_ = "STUB: not implemented"
	return *new(NewPutOpenai)
}

func New(tp elastictransport.Interface) *PutOpenai { _ = "STUB: not implemented"; return nil }

func (r *PutOpenai) Raw(raw io.Reader) *PutOpenai { _ = "STUB: not implemented"; return nil }

func (r *PutOpenai) Request(req *Request) *PutOpenai { _ = "STUB: not implemented"; return nil }

func (r *PutOpenai) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutOpenai) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutOpenai) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutOpenai) Header(key, value string) *PutOpenai { _ = "STUB: not implemented"; return nil }

func (r *PutOpenai) _tasktype(tasktype string) *PutOpenai { _ = "STUB: not implemented"; return nil }

func (r *PutOpenai) _openaiinferenceid(openaiinferenceid string) *PutOpenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenai) Timeout(duration string) *PutOpenai { _ = "STUB: not implemented"; return nil }

func (r *PutOpenai) ErrorTrace(errortrace bool) *PutOpenai { _ = "STUB: not implemented"; return nil }

func (r *PutOpenai) FilterPath(filterpaths ...string) *PutOpenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenai) Human(human bool) *PutOpenai { _ = "STUB: not implemented"; return nil }

func (r *PutOpenai) Pretty(pretty bool) *PutOpenai { _ = "STUB: not implemented"; return nil }

func (r *PutOpenai) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutOpenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenai) Service(service openaiservicetype.OpenAIServiceType) *PutOpenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenai) ServiceSettings(servicesettings types.OpenAIServiceSettingsVariant) *PutOpenai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenai) TaskSettings(tasksettings types.OpenAITaskSettingsVariant) *PutOpenai {
	_ = "STUB: not implemented"
	return nil
}
