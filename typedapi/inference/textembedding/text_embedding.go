package textembedding

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

type TextEmbedding struct {
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

type NewTextEmbedding func(inferenceid string) *TextEmbedding

func NewTextEmbeddingFunc(tp elastictransport.Interface) NewTextEmbedding {
	_ = "STUB: not implemented"
	return *new(NewTextEmbedding)
}

func New(tp elastictransport.Interface) *TextEmbedding { _ = "STUB: not implemented"; return nil }

func (r *TextEmbedding) Raw(raw io.Reader) *TextEmbedding { _ = "STUB: not implemented"; return nil }

func (r *TextEmbedding) Request(req *Request) *TextEmbedding { _ = "STUB: not implemented"; return nil }

func (r *TextEmbedding) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r TextEmbedding) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r TextEmbedding) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TextEmbedding) Header(key, value string) *TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *TextEmbedding) _inferenceid(inferenceid string) *TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *TextEmbedding) Timeout(duration string) *TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *TextEmbedding) ErrorTrace(errortrace bool) *TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *TextEmbedding) FilterPath(filterpaths ...string) *TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *TextEmbedding) Human(human bool) *TextEmbedding { _ = "STUB: not implemented"; return nil }

func (r *TextEmbedding) Pretty(pretty bool) *TextEmbedding { _ = "STUB: not implemented"; return nil }

func (r *TextEmbedding) Input(inputs ...string) *TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *TextEmbedding) InputType(inputtype string) *TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *TextEmbedding) TaskSettings(tasksettings json.RawMessage) *TextEmbedding {
	_ = "STUB: not implemented"
	return nil
}
