package previewdatafeed

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

type PreviewDatafeed struct {
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

type NewPreviewDatafeed func() *PreviewDatafeed

func NewPreviewDatafeedFunc(tp elastictransport.Interface) NewPreviewDatafeed {
	_ = "STUB: not implemented"
	return *new(NewPreviewDatafeed)
}

func New(tp elastictransport.Interface) *PreviewDatafeed { _ = "STUB: not implemented"; return nil }

func (r *PreviewDatafeed) Raw(raw io.Reader) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) Request(req *Request) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PreviewDatafeed) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PreviewDatafeed) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *PreviewDatafeed) Header(key, value string) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) DatafeedId(datafeedid string) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) Start(datetime string) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) End(datetime string) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) ErrorTrace(errortrace bool) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) FilterPath(filterpaths ...string) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) Human(human bool) *PreviewDatafeed { _ = "STUB: not implemented"; return nil }

func (r *PreviewDatafeed) Pretty(pretty bool) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) DatafeedConfig(datafeedconfig types.DatafeedConfigVariant) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PreviewDatafeed) JobConfig(jobconfig types.JobConfigVariant) *PreviewDatafeed {
	_ = "STUB: not implemented"
	return nil
}
