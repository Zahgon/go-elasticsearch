package putai21

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ai21servicetype"
)

const (
	tasktypeMask = iota + 1

	ai21inferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAi21 struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype        string
	ai21inferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutAi21 func(tasktype, ai21inferenceid string) *PutAi21

func NewPutAi21Func(tp elastictransport.Interface) NewPutAi21 {
	_ = "STUB: not implemented"
	return *new(NewPutAi21)
}

func New(tp elastictransport.Interface) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) Raw(raw io.Reader) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) Request(req *Request) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAi21) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAi21) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAi21) Header(key, value string) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) _tasktype(tasktype string) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) _ai21inferenceid(ai21inferenceid string) *PutAi21 {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAi21) Timeout(duration string) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) ErrorTrace(errortrace bool) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) FilterPath(filterpaths ...string) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) Human(human bool) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) Pretty(pretty bool) *PutAi21 { _ = "STUB: not implemented"; return nil }

func (r *PutAi21) Service(service ai21servicetype.Ai21ServiceType) *PutAi21 {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAi21) ServiceSettings(servicesettings types.Ai21ServiceSettingsVariant) *PutAi21 {
	_ = "STUB: not implemented"
	return nil
}
