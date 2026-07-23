package putdeepseek

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/deepseekservicetype"
)

const (
	tasktypeMask = iota + 1

	deepseekinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDeepseek struct {
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
	deepseekinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutDeepseek func(tasktype, deepseekinferenceid string) *PutDeepseek

func NewPutDeepseekFunc(tp elastictransport.Interface) NewPutDeepseek {
	_ = "STUB: not implemented"
	return *new(NewPutDeepseek)
}

func New(tp elastictransport.Interface) *PutDeepseek { _ = "STUB: not implemented"; return nil }

func (r *PutDeepseek) Raw(raw io.Reader) *PutDeepseek { _ = "STUB: not implemented"; return nil }

func (r *PutDeepseek) Request(req *Request) *PutDeepseek { _ = "STUB: not implemented"; return nil }

func (r *PutDeepseek) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDeepseek) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDeepseek) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutDeepseek) Header(key, value string) *PutDeepseek { _ = "STUB: not implemented"; return nil }

func (r *PutDeepseek) _tasktype(tasktype string) *PutDeepseek {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDeepseek) _deepseekinferenceid(deepseekinferenceid string) *PutDeepseek {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDeepseek) Timeout(duration string) *PutDeepseek { _ = "STUB: not implemented"; return nil }

func (r *PutDeepseek) ErrorTrace(errortrace bool) *PutDeepseek {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDeepseek) FilterPath(filterpaths ...string) *PutDeepseek {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDeepseek) Human(human bool) *PutDeepseek { _ = "STUB: not implemented"; return nil }

func (r *PutDeepseek) Pretty(pretty bool) *PutDeepseek { _ = "STUB: not implemented"; return nil }

func (r *PutDeepseek) Service(service deepseekservicetype.DeepSeekServiceType) *PutDeepseek {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDeepseek) ServiceSettings(servicesettings types.DeepSeekServiceSettingsVariant) *PutDeepseek {
	_ = "STUB: not implemented"
	return nil
}
