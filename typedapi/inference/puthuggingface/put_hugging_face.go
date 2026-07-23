package puthuggingface

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/huggingfaceservicetype"
)

const (
	tasktypeMask = iota + 1

	huggingfaceinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutHuggingFace struct {
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
	huggingfaceinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutHuggingFace func(tasktype, huggingfaceinferenceid string) *PutHuggingFace

func NewPutHuggingFaceFunc(tp elastictransport.Interface) NewPutHuggingFace {
	_ = "STUB: not implemented"
	return *new(NewPutHuggingFace)
}

func New(tp elastictransport.Interface) *PutHuggingFace { _ = "STUB: not implemented"; return nil }

func (r *PutHuggingFace) Raw(raw io.Reader) *PutHuggingFace { _ = "STUB: not implemented"; return nil }

func (r *PutHuggingFace) Request(req *Request) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutHuggingFace) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutHuggingFace) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutHuggingFace) Header(key, value string) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) _tasktype(tasktype string) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) _huggingfaceinferenceid(huggingfaceinferenceid string) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) Timeout(duration string) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) ErrorTrace(errortrace bool) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) FilterPath(filterpaths ...string) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) Human(human bool) *PutHuggingFace { _ = "STUB: not implemented"; return nil }

func (r *PutHuggingFace) Pretty(pretty bool) *PutHuggingFace { _ = "STUB: not implemented"; return nil }

func (r *PutHuggingFace) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) Service(service huggingfaceservicetype.HuggingFaceServiceType) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) ServiceSettings(servicesettings types.HuggingFaceServiceSettingsVariant) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutHuggingFace) TaskSettings(tasksettings types.HuggingFaceTaskSettingsVariant) *PutHuggingFace {
	_ = "STUB: not implemented"
	return nil
}
