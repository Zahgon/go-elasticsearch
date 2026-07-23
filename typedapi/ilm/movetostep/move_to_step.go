package movetostep

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
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MoveToStep struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewMoveToStep func(index string) *MoveToStep

func NewMoveToStepFunc(tp elastictransport.Interface) NewMoveToStep {
	_ = "STUB: not implemented"
	return *new(NewMoveToStep)
}

func New(tp elastictransport.Interface) *MoveToStep { _ = "STUB: not implemented"; return nil }

func (r *MoveToStep) Raw(raw io.Reader) *MoveToStep { _ = "STUB: not implemented"; return nil }

func (r *MoveToStep) Request(req *Request) *MoveToStep { _ = "STUB: not implemented"; return nil }

func (r *MoveToStep) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MoveToStep) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MoveToStep) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *MoveToStep) Header(key, value string) *MoveToStep { _ = "STUB: not implemented"; return nil }

func (r *MoveToStep) _index(index string) *MoveToStep { _ = "STUB: not implemented"; return nil }

func (r *MoveToStep) ErrorTrace(errortrace bool) *MoveToStep { _ = "STUB: not implemented"; return nil }

func (r *MoveToStep) FilterPath(filterpaths ...string) *MoveToStep {
	_ = "STUB: not implemented"
	return nil
}

func (r *MoveToStep) Human(human bool) *MoveToStep { _ = "STUB: not implemented"; return nil }

func (r *MoveToStep) Pretty(pretty bool) *MoveToStep { _ = "STUB: not implemented"; return nil }

func (r *MoveToStep) CurrentStep(currentstep types.StepKeyVariant) *MoveToStep {
	_ = "STUB: not implemented"
	return nil
}

func (r *MoveToStep) NextStep(nextstep types.StepKeyVariant) *MoveToStep {
	_ = "STUB: not implemented"
	return nil
}
