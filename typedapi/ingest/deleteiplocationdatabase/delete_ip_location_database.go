package deleteiplocationdatabase

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

type DeleteIpLocationDatabase struct {
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

type NewDeleteIpLocationDatabase func(id string) *DeleteIpLocationDatabase

func NewDeleteIpLocationDatabaseFunc(tp elastictransport.Interface) NewDeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return *new(NewDeleteIpLocationDatabase)
}

func New(tp elastictransport.Interface) *DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIpLocationDatabase) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteIpLocationDatabase) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteIpLocationDatabase) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteIpLocationDatabase) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteIpLocationDatabase) Header(key, value string) *DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIpLocationDatabase) _id(id string) *DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIpLocationDatabase) MasterTimeout(duration string) *DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIpLocationDatabase) Timeout(duration string) *DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIpLocationDatabase) ErrorTrace(errortrace bool) *DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIpLocationDatabase) FilterPath(filterpaths ...string) *DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIpLocationDatabase) Human(human bool) *DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIpLocationDatabase) Pretty(pretty bool) *DeleteIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}
