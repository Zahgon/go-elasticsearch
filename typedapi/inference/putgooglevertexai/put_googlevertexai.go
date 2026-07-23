package putgooglevertexai

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/googlevertexaiservicetype"
)

const (
	tasktypeMask = iota + 1

	googlevertexaiinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutGooglevertexai struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                  string
	googlevertexaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutGooglevertexai func(tasktype, googlevertexaiinferenceid string) *PutGooglevertexai

func NewPutGooglevertexaiFunc(tp elastictransport.Interface) NewPutGooglevertexai {
	_ = "STUB: not implemented"
	return *new(NewPutGooglevertexai)
}

func New(tp elastictransport.Interface) *PutGooglevertexai { _ = "STUB: not implemented"; return nil }

func (r *PutGooglevertexai) Raw(raw io.Reader) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) Request(req *Request) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutGooglevertexai) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutGooglevertexai) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutGooglevertexai) Header(key, value string) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) _tasktype(tasktype string) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) _googlevertexaiinferenceid(googlevertexaiinferenceid string) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) Timeout(duration string) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) ErrorTrace(errortrace bool) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) FilterPath(filterpaths ...string) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) Human(human bool) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) Pretty(pretty bool) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) Service(service googlevertexaiservicetype.GoogleVertexAIServiceType) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) ServiceSettings(servicesettings types.GoogleVertexAIServiceSettingsVariant) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGooglevertexai) TaskSettings(tasksettings types.GoogleVertexAITaskSettingsVariant) *PutGooglevertexai {
	_ = "STUB: not implemented"
	return nil
}
