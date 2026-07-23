package syncjobget

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

type SyncJobGet struct {
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

type NewSyncJobGet func(connectorsyncjobid string) *SyncJobGet

func NewSyncJobGetFunc(tp elastictransport.Interface) NewSyncJobGet {
	_ = "STUB: not implemented"
	return *new(NewSyncJobGet)
}

func New(tp elastictransport.Interface) *SyncJobGet { _ = "STUB: not implemented"; return nil }

func (r *SyncJobGet) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobGet) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobGet) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobGet) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SyncJobGet) Header(key, value string) *SyncJobGet { _ = "STUB: not implemented"; return nil }

func (r *SyncJobGet) _connectorsyncjobid(connectorsyncjobid string) *SyncJobGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobGet) ErrorTrace(errortrace bool) *SyncJobGet { _ = "STUB: not implemented"; return nil }

func (r *SyncJobGet) FilterPath(filterpaths ...string) *SyncJobGet {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobGet) Human(human bool) *SyncJobGet { _ = "STUB: not implemented"; return nil }

func (r *SyncJobGet) Pretty(pretty bool) *SyncJobGet { _ = "STUB: not implemented"; return nil }
