package getsynonym

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSynonym struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetSynonym func(id string) *GetSynonym

func NewGetSynonymFunc(tp elastictransport.Interface) NewGetSynonym {
	_ = "STUB: not implemented"
	return *new(NewGetSynonym)
}

func New(tp elastictransport.Interface) *GetSynonym { _ = "STUB: not implemented"; return nil }

func (r *GetSynonym) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSynonym) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSynonym) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSynonym) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetSynonym) Header(key, value string) *GetSynonym { _ = "STUB: not implemented"; return nil }

func (r *GetSynonym) _id(id string) *GetSynonym { _ = "STUB: not implemented"; return nil }

func (r *GetSynonym) From(from int) *GetSynonym { _ = "STUB: not implemented"; return nil }

func (r *GetSynonym) Size(size int) *GetSynonym { _ = "STUB: not implemented"; return nil }

func (r *GetSynonym) SearchAfter(searchafter string) *GetSynonym {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonym) ErrorTrace(errortrace bool) *GetSynonym { _ = "STUB: not implemented"; return nil }

func (r *GetSynonym) FilterPath(filterpaths ...string) *GetSynonym {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonym) Human(human bool) *GetSynonym { _ = "STUB: not implemented"; return nil }

func (r *GetSynonym) Pretty(pretty bool) *GetSynonym { _ = "STUB: not implemented"; return nil }
