package getgeoipdatabase

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

type GetGeoipDatabase struct {
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

type NewGetGeoipDatabase func() *GetGeoipDatabase

func NewGetGeoipDatabaseFunc(tp elastictransport.Interface) NewGetGeoipDatabase {
	_ = "STUB: not implemented"
	return *new(NewGetGeoipDatabase)
}

func New(tp elastictransport.Interface) *GetGeoipDatabase { _ = "STUB: not implemented"; return nil }

func (r *GetGeoipDatabase) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetGeoipDatabase) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetGeoipDatabase) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetGeoipDatabase) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetGeoipDatabase) Header(key, value string) *GetGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetGeoipDatabase) Id(id string) *GetGeoipDatabase { _ = "STUB: not implemented"; return nil }

func (r *GetGeoipDatabase) ErrorTrace(errortrace bool) *GetGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetGeoipDatabase) FilterPath(filterpaths ...string) *GetGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetGeoipDatabase) Human(human bool) *GetGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetGeoipDatabase) Pretty(pretty bool) *GetGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}
