package clearrepositoriesmeteringarchive

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

	maxarchiveversionMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ClearRepositoriesMeteringArchive struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	nodeid            string
	maxarchiveversion string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewClearRepositoriesMeteringArchive func(nodeid, maxarchiveversion string) *ClearRepositoriesMeteringArchive

func NewClearRepositoriesMeteringArchiveFunc(tp elastictransport.Interface) NewClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return *new(NewClearRepositoriesMeteringArchive)
}

func New(tp elastictransport.Interface) *ClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearRepositoriesMeteringArchive) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearRepositoriesMeteringArchive) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearRepositoriesMeteringArchive) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearRepositoriesMeteringArchive) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ClearRepositoriesMeteringArchive) Header(key, value string) *ClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearRepositoriesMeteringArchive) _nodeid(nodeid string) *ClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearRepositoriesMeteringArchive) _maxarchiveversion(maxarchiveversion string) *ClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearRepositoriesMeteringArchive) ErrorTrace(errortrace bool) *ClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearRepositoriesMeteringArchive) FilterPath(filterpaths ...string) *ClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearRepositoriesMeteringArchive) Human(human bool) *ClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearRepositoriesMeteringArchive) Pretty(pretty bool) *ClearRepositoriesMeteringArchive {
	_ = "STUB: not implemented"
	return nil
}
