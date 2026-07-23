package deleteexpireddata

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
	jobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteExpiredData struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteExpiredData func() *DeleteExpiredData

func NewDeleteExpiredDataFunc(tp elastictransport.Interface) NewDeleteExpiredData {
	_ = "STUB: not implemented"
	return *new(NewDeleteExpiredData)
}

func New(tp elastictransport.Interface) *DeleteExpiredData { _ = "STUB: not implemented"; return nil }

func (r *DeleteExpiredData) Raw(raw io.Reader) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteExpiredData) Request(req *Request) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteExpiredData) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteExpiredData) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteExpiredData) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *DeleteExpiredData) Header(key, value string) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteExpiredData) JobId(jobid string) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteExpiredData) ErrorTrace(errortrace bool) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteExpiredData) FilterPath(filterpaths ...string) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteExpiredData) Human(human bool) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteExpiredData) Pretty(pretty bool) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteExpiredData) RequestsPerSecond(requestspersecond float32) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteExpiredData) Timeout(duration types.DurationVariant) *DeleteExpiredData {
	_ = "STUB: not implemented"
	return nil
}
