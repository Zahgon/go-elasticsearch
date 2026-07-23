package getrepositoriesmeteringinfo

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nodeidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetRepositoriesMeteringInfo struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	nodeid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetRepositoriesMeteringInfo func(nodeid string) *GetRepositoriesMeteringInfo

func NewGetRepositoriesMeteringInfoFunc(tp elastictransport.Interface) NewGetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return *new(NewGetRepositoriesMeteringInfo)
}

func New(tp elastictransport.Interface) *GetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepositoriesMeteringInfo) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRepositoriesMeteringInfo) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRepositoriesMeteringInfo) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRepositoriesMeteringInfo) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetRepositoriesMeteringInfo) Header(key, value string) *GetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepositoriesMeteringInfo) _nodeid(nodeid string) *GetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepositoriesMeteringInfo) ErrorTrace(errortrace bool) *GetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepositoriesMeteringInfo) FilterPath(filterpaths ...string) *GetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepositoriesMeteringInfo) Human(human bool) *GetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRepositoriesMeteringInfo) Pretty(pretty bool) *GetRepositoriesMeteringInfo {
	_ = "STUB: not implemented"
	return nil
}
