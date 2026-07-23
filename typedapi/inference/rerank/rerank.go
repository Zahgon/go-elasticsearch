package rerank

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	inferenceidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Rerank struct {
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

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRerank func(inferenceid string) *Rerank

func NewRerankFunc(tp elastictransport.Interface) NewRerank {
	_ = "STUB: not implemented"
	return *new(NewRerank)
}

func New(tp elastictransport.Interface) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) Raw(raw io.Reader) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) Request(req *Request) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Rerank) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Rerank) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *Rerank) Header(key, value string) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) _inferenceid(inferenceid string) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) Timeout(duration string) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) ErrorTrace(errortrace bool) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) FilterPath(filterpaths ...string) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) Human(human bool) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) Pretty(pretty bool) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) Input(inputs ...string) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) Query(query string) *Rerank { _ = "STUB: not implemented"; return nil }

func (r *Rerank) ReturnDocuments(returndocuments bool) *Rerank {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rerank) TaskSettings(tasksettings json.RawMessage) *Rerank {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rerank) TopN(topn int) *Rerank { _ = "STUB: not implemented"; return nil }
