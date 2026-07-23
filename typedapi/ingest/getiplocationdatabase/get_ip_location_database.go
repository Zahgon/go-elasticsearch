package getiplocationdatabase

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

type GetIpLocationDatabase struct {
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

type NewGetIpLocationDatabase func() *GetIpLocationDatabase

func NewGetIpLocationDatabaseFunc(tp elastictransport.Interface) NewGetIpLocationDatabase {
	_ = "STUB: not implemented"
	return *new(NewGetIpLocationDatabase)
}

func New(tp elastictransport.Interface) *GetIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIpLocationDatabase) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetIpLocationDatabase) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetIpLocationDatabase) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetIpLocationDatabase) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetIpLocationDatabase) Header(key, value string) *GetIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIpLocationDatabase) Id(id string) *GetIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIpLocationDatabase) ErrorTrace(errortrace bool) *GetIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIpLocationDatabase) FilterPath(filterpaths ...string) *GetIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIpLocationDatabase) Human(human bool) *GetIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIpLocationDatabase) Pretty(pretty bool) *GetIpLocationDatabase {
	_ = "STUB: not implemented"
	return nil
}
