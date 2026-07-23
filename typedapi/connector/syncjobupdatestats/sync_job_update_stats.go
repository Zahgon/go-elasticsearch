package syncjobupdatestats

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	connectorsyncjobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SyncJobUpdateStats struct {
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

type NewSyncJobUpdateStats func(connectorsyncjobid string) *SyncJobUpdateStats

func NewSyncJobUpdateStatsFunc(tp elastictransport.Interface) NewSyncJobUpdateStats {
	_ = "STUB: not implemented"
	return *new(NewSyncJobUpdateStats)
}

func New(tp elastictransport.Interface) *SyncJobUpdateStats { _ = "STUB: not implemented"; return nil }

func (r *SyncJobUpdateStats) Raw(raw io.Reader) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) Request(req *Request) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobUpdateStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobUpdateStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SyncJobUpdateStats) Header(key, value string) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) _connectorsyncjobid(connectorsyncjobid string) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) ErrorTrace(errortrace bool) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) FilterPath(filterpaths ...string) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) Human(human bool) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) Pretty(pretty bool) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) DeletedDocumentCount(deleteddocumentcount int64) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) IndexedDocumentCount(indexeddocumentcount int64) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) IndexedDocumentVolume(indexeddocumentvolume int64) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) LastSeen(duration types.DurationVariant) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) Metadata(metadata types.MetadataVariant) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobUpdateStats) TotalDocumentCount(totaldocumentcount int) *SyncJobUpdateStats {
	_ = "STUB: not implemented"
	return nil
}
