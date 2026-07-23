package getsynonymssets

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSynonymsSets struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetSynonymsSets func() *GetSynonymsSets

func NewGetSynonymsSetsFunc(tp elastictransport.Interface) NewGetSynonymsSets {
	_ = "STUB: not implemented"
	return *new(NewGetSynonymsSets)
}

func New(tp elastictransport.Interface) *GetSynonymsSets { _ = "STUB: not implemented"; return nil }

func (r *GetSynonymsSets) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSynonymsSets) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSynonymsSets) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSynonymsSets) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetSynonymsSets) Header(key, value string) *GetSynonymsSets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonymsSets) From(from int) *GetSynonymsSets { _ = "STUB: not implemented"; return nil }

func (r *GetSynonymsSets) Size(size int) *GetSynonymsSets { _ = "STUB: not implemented"; return nil }

func (r *GetSynonymsSets) ErrorTrace(errortrace bool) *GetSynonymsSets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonymsSets) FilterPath(filterpaths ...string) *GetSynonymsSets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonymsSets) Human(human bool) *GetSynonymsSets { _ = "STUB: not implemented"; return nil }

func (r *GetSynonymsSets) Pretty(pretty bool) *GetSynonymsSets {
	_ = "STUB: not implemented"
	return nil
}
