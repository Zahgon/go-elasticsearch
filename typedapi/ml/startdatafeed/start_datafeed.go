package startdatafeed

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

const (
	datafeedidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StartDatafeed struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	datafeedid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewStartDatafeed func(datafeedid string) *StartDatafeed

func NewStartDatafeedFunc(tp elastictransport.Interface) NewStartDatafeed {
	_ = "STUB: not implemented"
	return *new(NewStartDatafeed)
}

func New(tp elastictransport.Interface) *StartDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StartDatafeed) Raw(raw io.Reader) *StartDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StartDatafeed) Request(req *Request) *StartDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StartDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartDatafeed) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartDatafeed) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *StartDatafeed) Header(key, value string) *StartDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDatafeed) _datafeedid(datafeedid string) *StartDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDatafeed) ErrorTrace(errortrace bool) *StartDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDatafeed) FilterPath(filterpaths ...string) *StartDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDatafeed) Human(human bool) *StartDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StartDatafeed) Pretty(pretty bool) *StartDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StartDatafeed) End(datetime types.DateTimeVariant) *StartDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDatafeed) Start(datetime types.DateTimeVariant) *StartDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartDatafeed) Timeout(duration types.DurationVariant) *StartDatafeed {
	_ = "STUB: not implemented"
	return nil
}
