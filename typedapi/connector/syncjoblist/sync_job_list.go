package syncjoblist

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncjobtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncstatus"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SyncJobList struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSyncJobList func() *SyncJobList

func NewSyncJobListFunc(tp elastictransport.Interface) NewSyncJobList {
	_ = "STUB: not implemented"
	return *new(NewSyncJobList)
}

func New(tp elastictransport.Interface) *SyncJobList { _ = "STUB: not implemented"; return nil }

func (r *SyncJobList) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobList) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobList) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobList) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SyncJobList) Header(key, value string) *SyncJobList { _ = "STUB: not implemented"; return nil }

func (r *SyncJobList) From(from int) *SyncJobList { _ = "STUB: not implemented"; return nil }

func (r *SyncJobList) Size(size int) *SyncJobList { _ = "STUB: not implemented"; return nil }

func (r *SyncJobList) Status(status syncstatus.SyncStatus) *SyncJobList {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobList) ConnectorId(id string) *SyncJobList { _ = "STUB: not implemented"; return nil }

func (r *SyncJobList) JobType(jobtypes ...syncjobtype.SyncJobType) *SyncJobList {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobList) ErrorTrace(errortrace bool) *SyncJobList {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobList) FilterPath(filterpaths ...string) *SyncJobList {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobList) Human(human bool) *SyncJobList { _ = "STUB: not implemented"; return nil }

func (r *SyncJobList) Pretty(pretty bool) *SyncJobList { _ = "STUB: not implemented"; return nil }
