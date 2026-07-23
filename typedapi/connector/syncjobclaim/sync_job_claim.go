package syncjobclaim

import (
	gobytes "bytes"
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

type SyncJobClaim struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorsyncjobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSyncJobClaim func(connectorsyncjobid string) *SyncJobClaim

func NewSyncJobClaimFunc(tp elastictransport.Interface) NewSyncJobClaim {
	_ = "STUB: not implemented"
	return *new(NewSyncJobClaim)
}

func New(tp elastictransport.Interface) *SyncJobClaim { _ = "STUB: not implemented"; return nil }

func (r *SyncJobClaim) Raw(raw io.Reader) *SyncJobClaim { _ = "STUB: not implemented"; return nil }

func (r *SyncJobClaim) Request(req *Request) *SyncJobClaim { _ = "STUB: not implemented"; return nil }

func (r *SyncJobClaim) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobClaim) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobClaim) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SyncJobClaim) Header(key, value string) *SyncJobClaim {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobClaim) _connectorsyncjobid(connectorsyncjobid string) *SyncJobClaim {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobClaim) ErrorTrace(errortrace bool) *SyncJobClaim {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobClaim) FilterPath(filterpaths ...string) *SyncJobClaim {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobClaim) Human(human bool) *SyncJobClaim { _ = "STUB: not implemented"; return nil }

func (r *SyncJobClaim) Pretty(pretty bool) *SyncJobClaim { _ = "STUB: not implemented"; return nil }

func (r *SyncJobClaim) SyncCursor(synccursor any) *SyncJobClaim {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobClaim) WorkerHostname(workerhostname string) *SyncJobClaim {
	_ = "STUB: not implemented"
	return nil
}
