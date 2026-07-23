package syncjobpost

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncjobtriggermethod"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/syncjobtype"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SyncJobPost struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSyncJobPost func() *SyncJobPost

func NewSyncJobPostFunc(tp elastictransport.Interface) NewSyncJobPost {
	_ = "STUB: not implemented"
	return *new(NewSyncJobPost)
}

func New(tp elastictransport.Interface) *SyncJobPost { _ = "STUB: not implemented"; return nil }

func (r *SyncJobPost) Raw(raw io.Reader) *SyncJobPost { _ = "STUB: not implemented"; return nil }

func (r *SyncJobPost) Request(req *Request) *SyncJobPost { _ = "STUB: not implemented"; return nil }

func (r *SyncJobPost) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobPost) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SyncJobPost) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SyncJobPost) Header(key, value string) *SyncJobPost { _ = "STUB: not implemented"; return nil }

func (r *SyncJobPost) ErrorTrace(errortrace bool) *SyncJobPost {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobPost) FilterPath(filterpaths ...string) *SyncJobPost {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobPost) Human(human bool) *SyncJobPost { _ = "STUB: not implemented"; return nil }

func (r *SyncJobPost) Pretty(pretty bool) *SyncJobPost { _ = "STUB: not implemented"; return nil }

func (r *SyncJobPost) Id(id string) *SyncJobPost { _ = "STUB: not implemented"; return nil }

func (r *SyncJobPost) JobType(jobtype syncjobtype.SyncJobType) *SyncJobPost {
	_ = "STUB: not implemented"
	return nil
}

func (r *SyncJobPost) TriggerMethod(triggermethod syncjobtriggermethod.SyncJobTriggerMethod) *SyncJobPost {
	_ = "STUB: not implemented"
	return nil
}
