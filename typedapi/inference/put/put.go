package put

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	tasktypeMask = iota + 1

	inferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Put struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype    string
	inferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPut func(inferenceid string) *Put

func NewPutFunc(tp elastictransport.Interface) NewPut {
	_ = "STUB: not implemented"
	return *new(NewPut)
}

func New(tp elastictransport.Interface) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) Raw(raw io.Reader) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) Request(req *Request) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Put) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Put) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Put) Header(key, value string) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) TaskType(tasktype string) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) _inferenceid(inferenceid string) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) Timeout(duration string) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) ErrorTrace(errortrace bool) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) FilterPath(filterpaths ...string) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) Human(human bool) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) Pretty(pretty bool) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *Put {
	_ = "STUB: not implemented"
	return nil
}

func (r *Put) Service(service string) *Put { _ = "STUB: not implemented"; return nil }

func (r *Put) ServiceSettings(servicesettings json.RawMessage) *Put {
	_ = "STUB: not implemented"
	return nil
}

func (r *Put) TaskSettings(tasksettings json.RawMessage) *Put {
	_ = "STUB: not implemented"
	return nil
}
