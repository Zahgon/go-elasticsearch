package translate

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Translate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewTranslate func() *Translate

func NewTranslateFunc(tp elastictransport.Interface) NewTranslate {
	_ = "STUB: not implemented"
	return *new(NewTranslate)
}

func New(tp elastictransport.Interface) *Translate { _ = "STUB: not implemented"; return nil }

func (r *Translate) Raw(raw io.Reader) *Translate { _ = "STUB: not implemented"; return nil }

func (r *Translate) Request(req *Request) *Translate { _ = "STUB: not implemented"; return nil }

func (r *Translate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Translate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Translate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Translate) Header(key, value string) *Translate { _ = "STUB: not implemented"; return nil }

func (r *Translate) ErrorTrace(errortrace bool) *Translate { _ = "STUB: not implemented"; return nil }

func (r *Translate) FilterPath(filterpaths ...string) *Translate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Translate) Human(human bool) *Translate { _ = "STUB: not implemented"; return nil }

func (r *Translate) Pretty(pretty bool) *Translate { _ = "STUB: not implemented"; return nil }

func (r *Translate) FetchSize(fetchsize int) *Translate { _ = "STUB: not implemented"; return nil }

func (r *Translate) Filter(filter types.QueryVariant) *Translate {
	_ = "STUB: not implemented"
	return nil
}

func (r *Translate) Query(query string) *Translate { _ = "STUB: not implemented"; return nil }

func (r *Translate) TimeZone(timezone string) *Translate { _ = "STUB: not implemented"; return nil }
