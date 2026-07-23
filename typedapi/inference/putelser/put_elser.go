package putelser

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/elserservicetype"
)

const (
	tasktypeMask = iota + 1

	elserinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutElser struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype         string
	elserinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutElser func(tasktype, elserinferenceid string) *PutElser

func NewPutElserFunc(tp elastictransport.Interface) NewPutElser {
	_ = "STUB: not implemented"
	return *new(NewPutElser)
}

func New(tp elastictransport.Interface) *PutElser { _ = "STUB: not implemented"; return nil }

func (r *PutElser) Raw(raw io.Reader) *PutElser { _ = "STUB: not implemented"; return nil }

func (r *PutElser) Request(req *Request) *PutElser { _ = "STUB: not implemented"; return nil }

func (r *PutElser) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutElser) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutElser) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutElser) Header(key, value string) *PutElser { _ = "STUB: not implemented"; return nil }

func (r *PutElser) _tasktype(tasktype string) *PutElser { _ = "STUB: not implemented"; return nil }

func (r *PutElser) _elserinferenceid(elserinferenceid string) *PutElser {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElser) Timeout(duration string) *PutElser { _ = "STUB: not implemented"; return nil }

func (r *PutElser) ErrorTrace(errortrace bool) *PutElser { _ = "STUB: not implemented"; return nil }

func (r *PutElser) FilterPath(filterpaths ...string) *PutElser {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElser) Human(human bool) *PutElser { _ = "STUB: not implemented"; return nil }

func (r *PutElser) Pretty(pretty bool) *PutElser { _ = "STUB: not implemented"; return nil }

func (r *PutElser) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutElser {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElser) Service(service elserservicetype.ElserServiceType) *PutElser {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutElser) ServiceSettings(servicesettings types.ElserServiceSettingsVariant) *PutElser {
	_ = "STUB: not implemented"
	return nil
}
