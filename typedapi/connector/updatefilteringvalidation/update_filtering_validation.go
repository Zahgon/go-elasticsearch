package updatefilteringvalidation

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
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateFilteringValidation struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdateFilteringValidation func(connectorid string) *UpdateFilteringValidation

func NewUpdateFilteringValidationFunc(tp elastictransport.Interface) NewUpdateFilteringValidation {
	_ = "STUB: not implemented"
	return *new(NewUpdateFilteringValidation)
}

func New(tp elastictransport.Interface) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilteringValidation) Raw(raw io.Reader) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilteringValidation) Request(req *Request) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilteringValidation) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateFilteringValidation) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateFilteringValidation) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateFilteringValidation) Header(key, value string) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilteringValidation) _connectorid(connectorid string) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilteringValidation) ErrorTrace(errortrace bool) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilteringValidation) FilterPath(filterpaths ...string) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilteringValidation) Human(human bool) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilteringValidation) Pretty(pretty bool) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFilteringValidation) Validation(validation types.FilteringRulesValidationVariant) *UpdateFilteringValidation {
	_ = "STUB: not implemented"
	return nil
}
