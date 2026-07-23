package putgroq

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/groqservicetype"
)

const (
	tasktypeMask = iota + 1

	groqinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutGroq struct {
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
	groqinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutGroq func(tasktype, groqinferenceid string) *PutGroq

func NewPutGroqFunc(tp elastictransport.Interface) NewPutGroq {
	_ = "STUB: not implemented"
	return *new(NewPutGroq)
}

func New(tp elastictransport.Interface) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) Raw(raw io.Reader) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) Request(req *Request) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutGroq) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutGroq) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutGroq) Header(key, value string) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) _tasktype(tasktype string) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) _groqinferenceid(groqinferenceid string) *PutGroq {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGroq) Timeout(duration string) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) ErrorTrace(errortrace bool) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) FilterPath(filterpaths ...string) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) Human(human bool) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) Pretty(pretty bool) *PutGroq { _ = "STUB: not implemented"; return nil }

func (r *PutGroq) Service(service groqservicetype.GroqServiceType) *PutGroq {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGroq) ServiceSettings(servicesettings types.GroqServiceSettingsVariant) *PutGroq {
	_ = "STUB: not implemented"
	return nil
}
