package delete

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Delete struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDelete func(name string) *Delete

func NewDeleteFunc(tp elastictransport.Interface) NewDelete {
	_ = "STUB: not implemented"
	return *new(NewDelete)
}

func New(tp elastictransport.Interface) *Delete { _ = "STUB: not implemented"; return nil }

func (r *Delete) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Delete) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Delete) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Delete) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Delete) Header(key, value string) *Delete { _ = "STUB: not implemented"; return nil }

func (r *Delete) _name(name string) *Delete { _ = "STUB: not implemented"; return nil }

func (r *Delete) ErrorTrace(errortrace bool) *Delete { _ = "STUB: not implemented"; return nil }

func (r *Delete) FilterPath(filterpaths ...string) *Delete { _ = "STUB: not implemented"; return nil }

func (r *Delete) Human(human bool) *Delete { _ = "STUB: not implemented"; return nil }

func (r *Delete) Pretty(pretty bool) *Delete { _ = "STUB: not implemented"; return nil }
