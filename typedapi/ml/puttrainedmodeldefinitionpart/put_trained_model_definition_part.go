package puttrainedmodeldefinitionpart

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	modelidMask = iota + 1

	partMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutTrainedModelDefinitionPart struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	modelid string
	part    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutTrainedModelDefinitionPart func(modelid, part string) *PutTrainedModelDefinitionPart

func NewPutTrainedModelDefinitionPartFunc(tp elastictransport.Interface) NewPutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return *new(NewPutTrainedModelDefinitionPart)
}

func New(tp elastictransport.Interface) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) Raw(raw io.Reader) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) Request(req *Request) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTrainedModelDefinitionPart) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTrainedModelDefinitionPart) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutTrainedModelDefinitionPart) Header(key, value string) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) _modelid(modelid string) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) _part(part string) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) ErrorTrace(errortrace bool) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) FilterPath(filterpaths ...string) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) Human(human bool) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) Pretty(pretty bool) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) Definition(definition string) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) TotalDefinitionLength(totaldefinitionlength int64) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelDefinitionPart) TotalParts(totalparts int) *PutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return nil
}
