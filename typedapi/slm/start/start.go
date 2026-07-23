package start

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Start struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewStart func() *Start

func NewStartFunc(tp elastictransport.Interface) NewStart {
	_ = "STUB: not implemented"
	return *new(NewStart)
}

func New(tp elastictransport.Interface) *Start { _ = "STUB: not implemented"; return nil }

func (r *Start) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Start) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Start) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Start) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Start) Header(key, value string) *Start { _ = "STUB: not implemented"; return nil }

func (r *Start) MasterTimeout(duration string) *Start { _ = "STUB: not implemented"; return nil }

func (r *Start) Timeout(duration string) *Start { _ = "STUB: not implemented"; return nil }

func (r *Start) ErrorTrace(errortrace bool) *Start { _ = "STUB: not implemented"; return nil }

func (r *Start) FilterPath(filterpaths ...string) *Start { _ = "STUB: not implemented"; return nil }

func (r *Start) Human(human bool) *Start { _ = "STUB: not implemented"; return nil }

func (r *Start) Pretty(pretty bool) *Start { _ = "STUB: not implemented"; return nil }
