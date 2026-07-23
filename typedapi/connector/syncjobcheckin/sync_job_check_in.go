package syncjobcheckin

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

type SyncJobCheckIn struct {
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

type NewSyncJobCheckIn func(connectorsyncjobid string) *SyncJobCheckIn

func NewSyncJobCheckInFunc(tp elastictransport.Interface) NewSyncJobCheckIn {
	_ = "STUB: not implemented"
	return *new(NewSyncJobCheckIn)
}

func New(tp elastictransport.Interface) *SyncJobCheckIn { _ = "STUB: not implemented"; return nil }

func (r *SyncJobCheckIn) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobCheckIn) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobCheckIn) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobCheckIn) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SyncJobCheckIn) Header(key, value string) *SyncJobCheckIn {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobCheckIn) _connectorsyncjobid(connectorsyncjobid string) *SyncJobCheckIn {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobCheckIn) ErrorTrace(errortrace bool) *SyncJobCheckIn {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobCheckIn) FilterPath(filterpaths ...string) *SyncJobCheckIn {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobCheckIn) Human(human bool) *SyncJobCheckIn { _ = "STUB: not implemented"; return nil }

func (r *SyncJobCheckIn) Pretty(pretty bool) *SyncJobCheckIn { _ = "STUB: not implemented"; return nil }
