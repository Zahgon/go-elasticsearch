package deletegeoipdatabase

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

type DeleteGeoipDatabase struct {
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

type NewDeleteGeoipDatabase func(id string) *DeleteGeoipDatabase

func NewDeleteGeoipDatabaseFunc(tp elastictransport.Interface) NewDeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return *new(NewDeleteGeoipDatabase)
}

func New(tp elastictransport.Interface) *DeleteGeoipDatabase { _ = "STUB: not implemented"; return nil }

func (r *DeleteGeoipDatabase) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteGeoipDatabase) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteGeoipDatabase) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteGeoipDatabase) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteGeoipDatabase) Header(key, value string) *DeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteGeoipDatabase) _id(id string) *DeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteGeoipDatabase) MasterTimeout(duration string) *DeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteGeoipDatabase) Timeout(duration string) *DeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteGeoipDatabase) ErrorTrace(errortrace bool) *DeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteGeoipDatabase) FilterPath(filterpaths ...string) *DeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteGeoipDatabase) Human(human bool) *DeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteGeoipDatabase) Pretty(pretty bool) *DeleteGeoipDatabase {
	_ = "STUB: not implemented"
	return nil
}
