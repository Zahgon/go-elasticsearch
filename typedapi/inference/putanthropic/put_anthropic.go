package putanthropic

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/anthropicservicetype"
)

const (
	tasktypeMask = iota + 1

	anthropicinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAnthropic struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype             string
	anthropicinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutAnthropic func(tasktype, anthropicinferenceid string) *PutAnthropic

func NewPutAnthropicFunc(tp elastictransport.Interface) NewPutAnthropic {
	_ = "STUB: not implemented"
	return *new(NewPutAnthropic)
}

func New(tp elastictransport.Interface) *PutAnthropic { _ = "STUB: not implemented"; return nil }

func (r *PutAnthropic) Raw(raw io.Reader) *PutAnthropic { _ = "STUB: not implemented"; return nil }

func (r *PutAnthropic) Request(req *Request) *PutAnthropic { _ = "STUB: not implemented"; return nil }

func (r *PutAnthropic) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAnthropic) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAnthropic) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAnthropic) Header(key, value string) *PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAnthropic) _tasktype(tasktype string) *PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAnthropic) _anthropicinferenceid(anthropicinferenceid string) *PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAnthropic) Timeout(duration string) *PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAnthropic) ErrorTrace(errortrace bool) *PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAnthropic) FilterPath(filterpaths ...string) *PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAnthropic) Human(human bool) *PutAnthropic { _ = "STUB: not implemented"; return nil }

func (r *PutAnthropic) Pretty(pretty bool) *PutAnthropic { _ = "STUB: not implemented"; return nil }

func (r *PutAnthropic) Service(service anthropicservicetype.AnthropicServiceType) *PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAnthropic) ServiceSettings(servicesettings types.AnthropicServiceSettingsVariant) *PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAnthropic) TaskSettings(tasksettings types.AnthropicTaskSettingsVariant) *PutAnthropic {
	_ = "STUB: not implemented"
	return nil
}
