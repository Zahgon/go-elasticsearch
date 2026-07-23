package sparseembedding

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

type SparseEmbedding struct {
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

type NewSparseEmbedding func(inferenceid string) *SparseEmbedding

func NewSparseEmbeddingFunc(tp elastictransport.Interface) NewSparseEmbedding {
	_ = "STUB: not implemented"
	return *new(NewSparseEmbedding)
}

func New(tp elastictransport.Interface) *SparseEmbedding { _ = "STUB: not implemented"; return nil }

func (r *SparseEmbedding) Raw(raw io.Reader) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *SparseEmbedding) Request(req *Request) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *SparseEmbedding) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SparseEmbedding) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SparseEmbedding) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *SparseEmbedding) Header(key, value string) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *SparseEmbedding) _inferenceid(inferenceid string) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *SparseEmbedding) Timeout(duration string) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *SparseEmbedding) ErrorTrace(errortrace bool) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *SparseEmbedding) FilterPath(filterpaths ...string) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *SparseEmbedding) Human(human bool) *SparseEmbedding { _ = "STUB: not implemented"; return nil }

func (r *SparseEmbedding) Pretty(pretty bool) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *SparseEmbedding) Input(inputs ...string) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *SparseEmbedding) TaskSettings(tasksettings json.RawMessage) *SparseEmbedding {
	_ = "STUB: not implemented"
	return nil
}
