package closepointintime

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ClosePointInTime struct {
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

type NewClosePointInTime func() *ClosePointInTime

func NewClosePointInTimeFunc(tp elastictransport.Interface) NewClosePointInTime {
	_ = "STUB: not implemented"
	return *new(NewClosePointInTime)
}

func New(tp elastictransport.Interface) *ClosePointInTime { _ = "STUB: not implemented"; return nil }

func (r *ClosePointInTime) Raw(raw io.Reader) *ClosePointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClosePointInTime) Request(req *Request) *ClosePointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClosePointInTime) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClosePointInTime) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClosePointInTime) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ClosePointInTime) Header(key, value string) *ClosePointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClosePointInTime) ErrorTrace(errortrace bool) *ClosePointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClosePointInTime) FilterPath(filterpaths ...string) *ClosePointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClosePointInTime) Human(human bool) *ClosePointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClosePointInTime) Pretty(pretty bool) *ClosePointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClosePointInTime) Id(id string) *ClosePointInTime { _ = "STUB: not implemented"; return nil }
