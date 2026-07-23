package existssource

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

type ExistsSource struct {
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

type NewExistsSource func(index, id string) *ExistsSource

func NewExistsSourceFunc(tp elastictransport.Interface) NewExistsSource {
	_ = "STUB: not implemented"
	return *new(NewExistsSource)
}

func New(tp elastictransport.Interface) *ExistsSource { _ = "STUB: not implemented"; return nil }

func (r *ExistsSource) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsSource) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsSource) Do(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r ExistsSource) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExistsSource) Header(key, value string) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) _id(id string) *ExistsSource { _ = "STUB: not implemented"; return nil }

func (r *ExistsSource) _index(index string) *ExistsSource { _ = "STUB: not implemented"; return nil }

func (r *ExistsSource) Preference(preference string) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) Realtime(realtime bool) *ExistsSource { _ = "STUB: not implemented"; return nil }

func (r *ExistsSource) Refresh(refresh bool) *ExistsSource { _ = "STUB: not implemented"; return nil }

func (r *ExistsSource) Routing(routings ...string) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) Source_(sourceconfigparam string) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) SourceExcludes_(fields ...string) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) SourceIncludes_(fields ...string) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) Version(versionnumber string) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) VersionType(versiontype versiontype.VersionType) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) ErrorTrace(errortrace bool) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) FilterPath(filterpaths ...string) *ExistsSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsSource) Human(human bool) *ExistsSource { _ = "STUB: not implemented"; return nil }

func (r *ExistsSource) Pretty(pretty bool) *ExistsSource { _ = "STUB: not implemented"; return nil }
