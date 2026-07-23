package stopdatafeed

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

type StopDatafeed struct {
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

type NewStopDatafeed func(datafeedid string) *StopDatafeed

func NewStopDatafeedFunc(tp elastictransport.Interface) NewStopDatafeed {
	_ = "STUB: not implemented"
	return *new(NewStopDatafeed)
}

func New(tp elastictransport.Interface) *StopDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StopDatafeed) Raw(raw io.Reader) *StopDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StopDatafeed) Request(req *Request) *StopDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StopDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopDatafeed) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopDatafeed) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *StopDatafeed) Header(key, value string) *StopDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDatafeed) _datafeedid(datafeedid string) *StopDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDatafeed) ErrorTrace(errortrace bool) *StopDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDatafeed) FilterPath(filterpaths ...string) *StopDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDatafeed) Human(human bool) *StopDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StopDatafeed) Pretty(pretty bool) *StopDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StopDatafeed) AllowNoMatch(allownomatch bool) *StopDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopDatafeed) CloseJob(closejob bool) *StopDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StopDatafeed) Force(force bool) *StopDatafeed { _ = "STUB: not implemented"; return nil }

func (r *StopDatafeed) Timeout(duration types.DurationVariant) *StopDatafeed {
	_ = "STUB: not implemented"
	return nil
}
