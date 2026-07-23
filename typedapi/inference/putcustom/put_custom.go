package putcustom

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/customservicetype"
)

const (
	tasktypeMask = iota + 1

	custominferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutCustom struct {
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
	custominferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutCustom func(tasktype, custominferenceid string) *PutCustom

func NewPutCustomFunc(tp elastictransport.Interface) NewPutCustom {
	_ = "STUB: not implemented"
	return *new(NewPutCustom)
}

func New(tp elastictransport.Interface) *PutCustom { _ = "STUB: not implemented"; return nil }

func (r *PutCustom) Raw(raw io.Reader) *PutCustom { _ = "STUB: not implemented"; return nil }

func (r *PutCustom) Request(req *Request) *PutCustom { _ = "STUB: not implemented"; return nil }

func (r *PutCustom) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutCustom) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutCustom) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutCustom) Header(key, value string) *PutCustom { _ = "STUB: not implemented"; return nil }

func (r *PutCustom) _tasktype(tasktype string) *PutCustom { _ = "STUB: not implemented"; return nil }

func (r *PutCustom) _custominferenceid(custominferenceid string) *PutCustom {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCustom) ErrorTrace(errortrace bool) *PutCustom { _ = "STUB: not implemented"; return nil }

func (r *PutCustom) FilterPath(filterpaths ...string) *PutCustom {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCustom) Human(human bool) *PutCustom { _ = "STUB: not implemented"; return nil }

func (r *PutCustom) Pretty(pretty bool) *PutCustom { _ = "STUB: not implemented"; return nil }

func (r *PutCustom) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutCustom {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCustom) Service(service customservicetype.CustomServiceType) *PutCustom {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCustom) ServiceSettings(servicesettings types.CustomServiceSettingsVariant) *PutCustom {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCustom) TaskSettings(tasksettings types.CustomTaskSettingsVariant) *PutCustom {
	_ = "STUB: not implemented"
	return nil
}
