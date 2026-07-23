package modifydatastream

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

type ModifyDataStream struct {
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

type NewModifyDataStream func() *ModifyDataStream

func NewModifyDataStreamFunc(tp elastictransport.Interface) NewModifyDataStream {
	_ = "STUB: not implemented"
	return *new(NewModifyDataStream)
}

func New(tp elastictransport.Interface) *ModifyDataStream { _ = "STUB: not implemented"; return nil }

func (r *ModifyDataStream) Raw(raw io.Reader) *ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *ModifyDataStream) Request(req *Request) *ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *ModifyDataStream) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ModifyDataStream) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ModifyDataStream) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ModifyDataStream) Header(key, value string) *ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *ModifyDataStream) ErrorTrace(errortrace bool) *ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *ModifyDataStream) FilterPath(filterpaths ...string) *ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *ModifyDataStream) Human(human bool) *ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *ModifyDataStream) Pretty(pretty bool) *ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *ModifyDataStream) Actions(actions ...types.IndicesModifyActionVariant) *ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}

func (r *ModifyDataStream) ActionsValues(actionsvalues []types.IndicesModifyAction) *ModifyDataStream {
	_ = "STUB: not implemented"
	return nil
}
