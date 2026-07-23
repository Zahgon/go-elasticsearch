package putjinaai

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaiservicetype"
)

const (
	tasktypeMask = iota + 1

	jinaaiinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutJinaai struct {
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
	jinaaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutJinaai func(tasktype, jinaaiinferenceid string) *PutJinaai

func NewPutJinaaiFunc(tp elastictransport.Interface) NewPutJinaai {
	_ = "STUB: not implemented"
	return *new(NewPutJinaai)
}

func New(tp elastictransport.Interface) *PutJinaai { _ = "STUB: not implemented"; return nil }

func (r *PutJinaai) Raw(raw io.Reader) *PutJinaai { _ = "STUB: not implemented"; return nil }

func (r *PutJinaai) Request(req *Request) *PutJinaai { _ = "STUB: not implemented"; return nil }

func (r *PutJinaai) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutJinaai) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutJinaai) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutJinaai) Header(key, value string) *PutJinaai { _ = "STUB: not implemented"; return nil }

func (r *PutJinaai) _tasktype(tasktype string) *PutJinaai { _ = "STUB: not implemented"; return nil }

func (r *PutJinaai) _jinaaiinferenceid(jinaaiinferenceid string) *PutJinaai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJinaai) Timeout(duration string) *PutJinaai { _ = "STUB: not implemented"; return nil }

func (r *PutJinaai) ErrorTrace(errortrace bool) *PutJinaai { _ = "STUB: not implemented"; return nil }

func (r *PutJinaai) FilterPath(filterpaths ...string) *PutJinaai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJinaai) Human(human bool) *PutJinaai { _ = "STUB: not implemented"; return nil }

func (r *PutJinaai) Pretty(pretty bool) *PutJinaai { _ = "STUB: not implemented"; return nil }

func (r *PutJinaai) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutJinaai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJinaai) Service(service jinaaiservicetype.JinaAIServiceType) *PutJinaai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJinaai) ServiceSettings(servicesettings types.JinaAIServiceSettingsVariant) *PutJinaai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJinaai) TaskSettings(tasksettings types.JinaAITaskSettingsVariant) *PutJinaai {
	_ = "STUB: not implemented"
	return nil
}
