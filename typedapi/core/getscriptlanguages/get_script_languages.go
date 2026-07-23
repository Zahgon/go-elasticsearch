package getscriptlanguages

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetScriptLanguages struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetScriptLanguages func() *GetScriptLanguages

func NewGetScriptLanguagesFunc(tp elastictransport.Interface) NewGetScriptLanguages {
	_ = "STUB: not implemented"
	return *new(NewGetScriptLanguages)
}

func New(tp elastictransport.Interface) *GetScriptLanguages { _ = "STUB: not implemented"; return nil }

func (r *GetScriptLanguages) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetScriptLanguages) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetScriptLanguages) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetScriptLanguages) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetScriptLanguages) Header(key, value string) *GetScriptLanguages {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScriptLanguages) ErrorTrace(errortrace bool) *GetScriptLanguages {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScriptLanguages) FilterPath(filterpaths ...string) *GetScriptLanguages {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScriptLanguages) Human(human bool) *GetScriptLanguages {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetScriptLanguages) Pretty(pretty bool) *GetScriptLanguages {
	_ = "STUB: not implemented"
	return nil
}
