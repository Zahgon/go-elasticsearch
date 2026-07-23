package deletesynonym

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

type DeleteSynonym struct {
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

type NewDeleteSynonym func(id string) *DeleteSynonym

func NewDeleteSynonymFunc(tp elastictransport.Interface) NewDeleteSynonym {
	_ = "STUB: not implemented"
	return *new(NewDeleteSynonym)
}

func New(tp elastictransport.Interface) *DeleteSynonym { _ = "STUB: not implemented"; return nil }

func (r *DeleteSynonym) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteSynonym) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteSynonym) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteSynonym) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteSynonym) Header(key, value string) *DeleteSynonym {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonym) _id(id string) *DeleteSynonym { _ = "STUB: not implemented"; return nil }

func (r *DeleteSynonym) ErrorTrace(errortrace bool) *DeleteSynonym {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonym) FilterPath(filterpaths ...string) *DeleteSynonym {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonym) Human(human bool) *DeleteSynonym { _ = "STUB: not implemented"; return nil }

func (r *DeleteSynonym) Pretty(pretty bool) *DeleteSynonym { _ = "STUB: not implemented"; return nil }
