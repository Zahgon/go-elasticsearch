package getsource

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSource struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id    string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetSource func(index, id string) *GetSource

func NewGetSourceFunc(tp elastictransport.Interface) NewGetSource {
	_ = "STUB: not implemented"
	return *new(NewGetSource)
}

func New(tp elastictransport.Interface) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSource) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSource) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetSource) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetSource) Header(key, value string) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) _id(id string) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) _index(index string) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) Preference(preference string) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) Realtime(realtime bool) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) Refresh(refresh bool) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) Routing(routings ...string) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) Source_(sourceconfigparam string) *GetSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSource) SourceExcludes_(fields ...string) *GetSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSource) SourceIncludes_(fields ...string) *GetSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSource) Version(versionnumber string) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) VersionType(versiontype versiontype.VersionType) *GetSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSource) ErrorTrace(errortrace bool) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) FilterPath(filterpaths ...string) *GetSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSource) Human(human bool) *GetSource { _ = "STUB: not implemented"; return nil }

func (r *GetSource) Pretty(pretty bool) *GetSource { _ = "STUB: not implemented"; return nil }
