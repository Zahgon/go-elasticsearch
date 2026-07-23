package putcontextualai

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/contextualaiservicetype"
)

const (
	tasktypeMask = iota + 1

	contextualaiinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutContextualai struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                string
	contextualaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutContextualai func(tasktype, contextualaiinferenceid string) *PutContextualai

func NewPutContextualaiFunc(tp elastictransport.Interface) NewPutContextualai {
	_ = "STUB: not implemented"
	return *new(NewPutContextualai)
}

func New(tp elastictransport.Interface) *PutContextualai { _ = "STUB: not implemented"; return nil }

func (r *PutContextualai) Raw(raw io.Reader) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) Request(req *Request) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutContextualai) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutContextualai) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutContextualai) Header(key, value string) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) _tasktype(tasktype string) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) _contextualaiinferenceid(contextualaiinferenceid string) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) Timeout(duration string) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) ErrorTrace(errortrace bool) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) FilterPath(filterpaths ...string) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) Human(human bool) *PutContextualai { _ = "STUB: not implemented"; return nil }

func (r *PutContextualai) Pretty(pretty bool) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) Service(service contextualaiservicetype.ContextualAIServiceType) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) ServiceSettings(servicesettings types.ContextualAIServiceSettingsVariant) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutContextualai) TaskSettings(tasksettings types.ContextualAITaskSettingsVariant) *PutContextualai {
	_ = "STUB: not implemented"
	return nil
}
