package putwatsonx

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/watsonxservicetype"
)

const (
	tasktypeMask = iota + 1

	watsonxinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutWatsonx struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype           string
	watsonxinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutWatsonx func(tasktype, watsonxinferenceid string) *PutWatsonx

func NewPutWatsonxFunc(tp elastictransport.Interface) NewPutWatsonx {
	_ = "STUB: not implemented"
	return *new(NewPutWatsonx)
}

func New(tp elastictransport.Interface) *PutWatsonx { _ = "STUB: not implemented"; return nil }

func (r *PutWatsonx) Raw(raw io.Reader) *PutWatsonx { _ = "STUB: not implemented"; return nil }

func (r *PutWatsonx) Request(req *Request) *PutWatsonx { _ = "STUB: not implemented"; return nil }

func (r *PutWatsonx) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutWatsonx) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutWatsonx) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutWatsonx) Header(key, value string) *PutWatsonx { _ = "STUB: not implemented"; return nil }

func (r *PutWatsonx) _tasktype(tasktype string) *PutWatsonx { _ = "STUB: not implemented"; return nil }

func (r *PutWatsonx) _watsonxinferenceid(watsonxinferenceid string) *PutWatsonx {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatsonx) Timeout(duration string) *PutWatsonx { _ = "STUB: not implemented"; return nil }

func (r *PutWatsonx) ErrorTrace(errortrace bool) *PutWatsonx { _ = "STUB: not implemented"; return nil }

func (r *PutWatsonx) FilterPath(filterpaths ...string) *PutWatsonx {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatsonx) Human(human bool) *PutWatsonx { _ = "STUB: not implemented"; return nil }

func (r *PutWatsonx) Pretty(pretty bool) *PutWatsonx { _ = "STUB: not implemented"; return nil }

func (r *PutWatsonx) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutWatsonx {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatsonx) Service(service watsonxservicetype.WatsonxServiceType) *PutWatsonx {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutWatsonx) ServiceSettings(servicesettings types.WatsonxServiceSettingsVariant) *PutWatsonx {
	_ = "STUB: not implemented"
	return nil
}
