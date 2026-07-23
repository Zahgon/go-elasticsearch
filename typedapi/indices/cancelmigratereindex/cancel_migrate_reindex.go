package cancelmigratereindex

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CancelMigrateReindex struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCancelMigrateReindex func(index string) *CancelMigrateReindex

func NewCancelMigrateReindexFunc(tp elastictransport.Interface) NewCancelMigrateReindex {
	_ = "STUB: not implemented"
	return *new(NewCancelMigrateReindex)
}

func New(tp elastictransport.Interface) *CancelMigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelMigrateReindex) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CancelMigrateReindex) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CancelMigrateReindex) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CancelMigrateReindex) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *CancelMigrateReindex) Header(key, value string) *CancelMigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelMigrateReindex) _index(index string) *CancelMigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelMigrateReindex) ErrorTrace(errortrace bool) *CancelMigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelMigrateReindex) FilterPath(filterpaths ...string) *CancelMigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelMigrateReindex) Human(human bool) *CancelMigrateReindex {
	_ = "STUB: not implemented"
	return nil
}

func (r *CancelMigrateReindex) Pretty(pretty bool) *CancelMigrateReindex {
	_ = "STUB: not implemented"
	return nil
}
