package syncjobdelete

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	connectorsyncjobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SyncJobDelete struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	connectorsyncjobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSyncJobDelete func(connectorsyncjobid string) *SyncJobDelete

func NewSyncJobDeleteFunc(tp elastictransport.Interface) NewSyncJobDelete {
	_ = "STUB: not implemented"
	return *new(NewSyncJobDelete)
}

func New(tp elastictransport.Interface) *SyncJobDelete { _ = "STUB: not implemented"; return nil }

func (r *SyncJobDelete) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobDelete) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobDelete) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobDelete) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SyncJobDelete) Header(key, value string) *SyncJobDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobDelete) _connectorsyncjobid(connectorsyncjobid string) *SyncJobDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobDelete) ErrorTrace(errortrace bool) *SyncJobDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobDelete) FilterPath(filterpaths ...string) *SyncJobDelete {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobDelete) Human(human bool) *SyncJobDelete { _ = "STUB: not implemented"; return nil }

func (r *SyncJobDelete) Pretty(pretty bool) *SyncJobDelete { _ = "STUB: not implemented"; return nil }
