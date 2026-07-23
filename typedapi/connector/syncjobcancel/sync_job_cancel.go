package syncjobcancel

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

type SyncJobCancel struct {
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

type NewSyncJobCancel func(connectorsyncjobid string) *SyncJobCancel

func NewSyncJobCancelFunc(tp elastictransport.Interface) NewSyncJobCancel {
	_ = "STUB: not implemented"
	return *new(NewSyncJobCancel)
}

func New(tp elastictransport.Interface) *SyncJobCancel { _ = "STUB: not implemented"; return nil }

func (r *SyncJobCancel) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobCancel) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobCancel) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobCancel) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SyncJobCancel) Header(key, value string) *SyncJobCancel {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobCancel) _connectorsyncjobid(connectorsyncjobid string) *SyncJobCancel {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobCancel) ErrorTrace(errortrace bool) *SyncJobCancel {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobCancel) FilterPath(filterpaths ...string) *SyncJobCancel {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobCancel) Human(human bool) *SyncJobCancel { _ = "STUB: not implemented"; return nil }

func (r *SyncJobCancel) Pretty(pretty bool) *SyncJobCancel { _ = "STUB: not implemented"; return nil }
