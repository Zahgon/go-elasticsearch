package getmigratereindexstatus

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

type GetMigrateReindexStatus struct {
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

type NewGetMigrateReindexStatus func(index string) *GetMigrateReindexStatus

func NewGetMigrateReindexStatusFunc(tp elastictransport.Interface) NewGetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return *new(NewGetMigrateReindexStatus)
}

func New(tp elastictransport.Interface) *GetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMigrateReindexStatus) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetMigrateReindexStatus) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetMigrateReindexStatus) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetMigrateReindexStatus) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetMigrateReindexStatus) Header(key, value string) *GetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMigrateReindexStatus) _index(index string) *GetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMigrateReindexStatus) ErrorTrace(errortrace bool) *GetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMigrateReindexStatus) FilterPath(filterpaths ...string) *GetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMigrateReindexStatus) Human(human bool) *GetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetMigrateReindexStatus) Pretty(pretty bool) *GetMigrateReindexStatus {
	_ = "STUB: not implemented"
	return nil
}
