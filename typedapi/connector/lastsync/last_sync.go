package lastsync

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncstatus"
)

const (
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type LastSync struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewLastSync func(connectorid string) *LastSync

func NewLastSyncFunc(tp elastictransport.Interface) NewLastSync {
	_ = "STUB: not implemented"
	return *new(NewLastSync)
}

func New(tp elastictransport.Interface) *LastSync { _ = "STUB: not implemented"; return nil }

func (r *LastSync) Raw(raw io.Reader) *LastSync { _ = "STUB: not implemented"; return nil }

func (r *LastSync) Request(req *Request) *LastSync { _ = "STUB: not implemented"; return nil }

func (r *LastSync) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r LastSync) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r LastSync) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *LastSync) Header(key, value string) *LastSync { _ = "STUB: not implemented"; return nil }

func (r *LastSync) _connectorid(connectorid string) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) ErrorTrace(errortrace bool) *LastSync { _ = "STUB: not implemented"; return nil }

func (r *LastSync) FilterPath(filterpaths ...string) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) Human(human bool) *LastSync { _ = "STUB: not implemented"; return nil }

func (r *LastSync) Pretty(pretty bool) *LastSync { _ = "STUB: not implemented"; return nil }

func (r *LastSync) LastAccessControlSyncError(lastaccesscontrolsyncerror string) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastAccessControlSyncScheduledAt(datetime types.DateTimeVariant) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastAccessControlSyncStatus(lastaccesscontrolsyncstatus syncstatus.SyncStatus) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastDeletedDocumentCount(lastdeleteddocumentcount int64) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastIncrementalSyncScheduledAt(datetime types.DateTimeVariant) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastIndexedDocumentCount(lastindexeddocumentcount int64) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastSeen(datetime types.DateTimeVariant) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastSyncError(lastsyncerror string) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastSyncScheduledAt(datetime types.DateTimeVariant) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastSyncStatus(lastsyncstatus syncstatus.SyncStatus) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) LastSynced(datetime types.DateTimeVariant) *LastSync {
	_ = "STUB: not implemented"
	return nil
}

func (r *LastSync) SyncCursor(synccursor any) *LastSync { _ = "STUB: not implemented"; return nil }
