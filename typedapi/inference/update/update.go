package update

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
	inferenceidMask = iota + 1

	tasktypeMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Update struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	inferenceid string
	tasktype    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdate func(inferenceid string) *Update

func NewUpdateFunc(tp elastictransport.Interface) NewUpdate {
	_ = "STUB: not implemented"
	return *new(NewUpdate)
}

func New(tp elastictransport.Interface) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Raw(raw io.Reader) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Request(req *Request) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Update) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Update) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Update) Header(key, value string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) _inferenceid(inferenceid string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) TaskType(tasktype string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Timeout(duration string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) ErrorTrace(errortrace bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) FilterPath(filterpaths ...string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Human(human bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) Pretty(pretty bool) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *Update {
	_ = "STUB: not implemented"
	return nil
}

func (r *Update) Service(service string) *Update { _ = "STUB: not implemented"; return nil }

func (r *Update) ServiceSettings(servicesettings json.RawMessage) *Update {
	_ = "STUB: not implemented"
	return nil
}

func (r *Update) TaskSettings(tasksettings json.RawMessage) *Update {
	_ = "STUB: not implemented"
	return nil
}
