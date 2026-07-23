package putmistral

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/mistralservicetype"
)

const (
	tasktypeMask = iota + 1

	mistralinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutMistral struct {
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
	mistralinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutMistral func(tasktype, mistralinferenceid string) *PutMistral

func NewPutMistralFunc(tp elastictransport.Interface) NewPutMistral {
	_ = "STUB: not implemented"
	return *new(NewPutMistral)
}

func New(tp elastictransport.Interface) *PutMistral { _ = "STUB: not implemented"; return nil }

func (r *PutMistral) Raw(raw io.Reader) *PutMistral { _ = "STUB: not implemented"; return nil }

func (r *PutMistral) Request(req *Request) *PutMistral { _ = "STUB: not implemented"; return nil }

func (r *PutMistral) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutMistral) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutMistral) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutMistral) Header(key, value string) *PutMistral { _ = "STUB: not implemented"; return nil }

func (r *PutMistral) _tasktype(tasktype string) *PutMistral { _ = "STUB: not implemented"; return nil }

func (r *PutMistral) _mistralinferenceid(mistralinferenceid string) *PutMistral {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMistral) Timeout(duration string) *PutMistral { _ = "STUB: not implemented"; return nil }

func (r *PutMistral) ErrorTrace(errortrace bool) *PutMistral { _ = "STUB: not implemented"; return nil }

func (r *PutMistral) FilterPath(filterpaths ...string) *PutMistral {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMistral) Human(human bool) *PutMistral { _ = "STUB: not implemented"; return nil }

func (r *PutMistral) Pretty(pretty bool) *PutMistral { _ = "STUB: not implemented"; return nil }

func (r *PutMistral) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutMistral {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMistral) Service(service mistralservicetype.MistralServiceType) *PutMistral {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutMistral) ServiceSettings(servicesettings types.MistralServiceSettingsVariant) *PutMistral {
	_ = "STUB: not implemented"
	return nil
}
