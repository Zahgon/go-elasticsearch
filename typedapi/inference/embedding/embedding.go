package embedding

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	inferenceidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Embedding struct {
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

type NewEmbedding func(inferenceid string) *Embedding

func NewEmbeddingFunc(tp elastictransport.Interface) NewEmbedding {
	_ = "STUB: not implemented"
	return *new(NewEmbedding)
}

func New(tp elastictransport.Interface) *Embedding { _ = "STUB: not implemented"; return nil }

func (r *Embedding) Raw(raw io.Reader) *Embedding { _ = "STUB: not implemented"; return nil }

func (r *Embedding) Request(req *Request) *Embedding { _ = "STUB: not implemented"; return nil }

func (r *Embedding) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Embedding) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Embedding) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Embedding) Header(key, value string) *Embedding { _ = "STUB: not implemented"; return nil }

func (r *Embedding) _inferenceid(inferenceid string) *Embedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *Embedding) Timeout(duration string) *Embedding { _ = "STUB: not implemented"; return nil }

func (r *Embedding) ErrorTrace(errortrace bool) *Embedding { _ = "STUB: not implemented"; return nil }

func (r *Embedding) FilterPath(filterpaths ...string) *Embedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *Embedding) Human(human bool) *Embedding { _ = "STUB: not implemented"; return nil }

func (r *Embedding) Pretty(pretty bool) *Embedding { _ = "STUB: not implemented"; return nil }

func (r *Embedding) Input(embeddinginput types.EmbeddingInputVariant) *Embedding {
	_ = "STUB: not implemented"
	return nil
}

func (r *Embedding) InputType(inputtype string) *Embedding { _ = "STUB: not implemented"; return nil }

func (r *Embedding) TaskSettings(tasksettings json.RawMessage) *Embedding {
	_ = "STUB: not implemented"
	return nil
}
