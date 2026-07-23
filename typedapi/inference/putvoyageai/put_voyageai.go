package putvoyageai

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/voyageaiservicetype"
)

const (
	tasktypeMask = iota + 1

	voyageaiinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutVoyageai struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype            string
	voyageaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutVoyageai func(tasktype, voyageaiinferenceid string) *PutVoyageai

func NewPutVoyageaiFunc(tp elastictransport.Interface) NewPutVoyageai {
	_ = "STUB: not implemented"
	return *new(NewPutVoyageai)
}

func New(tp elastictransport.Interface) *PutVoyageai { _ = "STUB: not implemented"; return nil }

func (r *PutVoyageai) Raw(raw io.Reader) *PutVoyageai { _ = "STUB: not implemented"; return nil }

func (r *PutVoyageai) Request(req *Request) *PutVoyageai { _ = "STUB: not implemented"; return nil }

func (r *PutVoyageai) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutVoyageai) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutVoyageai) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutVoyageai) Header(key, value string) *PutVoyageai { _ = "STUB: not implemented"; return nil }

func (r *PutVoyageai) _tasktype(tasktype string) *PutVoyageai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutVoyageai) _voyageaiinferenceid(voyageaiinferenceid string) *PutVoyageai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutVoyageai) Timeout(duration string) *PutVoyageai { _ = "STUB: not implemented"; return nil }

func (r *PutVoyageai) ErrorTrace(errortrace bool) *PutVoyageai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutVoyageai) FilterPath(filterpaths ...string) *PutVoyageai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutVoyageai) Human(human bool) *PutVoyageai { _ = "STUB: not implemented"; return nil }

func (r *PutVoyageai) Pretty(pretty bool) *PutVoyageai { _ = "STUB: not implemented"; return nil }

func (r *PutVoyageai) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutVoyageai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutVoyageai) Service(service voyageaiservicetype.VoyageAIServiceType) *PutVoyageai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutVoyageai) ServiceSettings(servicesettings types.VoyageAIServiceSettingsVariant) *PutVoyageai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutVoyageai) TaskSettings(tasksettings types.VoyageAITaskSettingsVariant) *PutVoyageai {
	_ = "STUB: not implemented"
	return nil
}
