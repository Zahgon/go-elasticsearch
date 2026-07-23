package putsynonym

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
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutSynonym struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutSynonym func(id string) *PutSynonym

func NewPutSynonymFunc(tp elastictransport.Interface) NewPutSynonym {
	_ = "STUB: not implemented"
	return *new(NewPutSynonym)
}

func New(tp elastictransport.Interface) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) Raw(raw io.Reader) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) Request(req *Request) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutSynonym) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutSynonym) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutSynonym) Header(key, value string) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) _id(id string) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) Refresh(refresh bool) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) Append(append bool) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) ErrorTrace(errortrace bool) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) FilterPath(filterpaths ...string) *PutSynonym {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSynonym) Human(human bool) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) Pretty(pretty bool) *PutSynonym { _ = "STUB: not implemented"; return nil }

func (r *PutSynonym) SynonymsSet(synonymssets ...types.SynonymRuleVariant) *PutSynonym {
	_ = "STUB: not implemented"
	return nil
}
