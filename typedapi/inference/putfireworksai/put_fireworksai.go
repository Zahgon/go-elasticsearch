package putfireworksai

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fireworksaiservicetype"
)

const (
	tasktypeMask = iota + 1

	fireworksaiinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutFireworksai struct {
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
	fireworksaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutFireworksai func(tasktype, fireworksaiinferenceid string) *PutFireworksai

func NewPutFireworksaiFunc(tp elastictransport.Interface) NewPutFireworksai {
	_ = "STUB: not implemented"
	return *new(NewPutFireworksai)
}

func New(tp elastictransport.Interface) *PutFireworksai { _ = "STUB: not implemented"; return nil }

func (r *PutFireworksai) Raw(raw io.Reader) *PutFireworksai { _ = "STUB: not implemented"; return nil }

func (r *PutFireworksai) Request(req *Request) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutFireworksai) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutFireworksai) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutFireworksai) Header(key, value string) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) _tasktype(tasktype string) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) _fireworksaiinferenceid(fireworksaiinferenceid string) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) Timeout(duration string) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) ErrorTrace(errortrace bool) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) FilterPath(filterpaths ...string) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) Human(human bool) *PutFireworksai { _ = "STUB: not implemented"; return nil }

func (r *PutFireworksai) Pretty(pretty bool) *PutFireworksai { _ = "STUB: not implemented"; return nil }

func (r *PutFireworksai) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) Service(service fireworksaiservicetype.FireworksAIServiceType) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) ServiceSettings(servicesettings types.FireworksAIServiceSettingsVariant) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutFireworksai) TaskSettings(tasksettings types.FireworksAITaskSettingsVariant) *PutFireworksai {
	_ = "STUB: not implemented"
	return nil
}
