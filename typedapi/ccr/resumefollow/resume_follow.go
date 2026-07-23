package resumefollow

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
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ResumeFollow struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewResumeFollow func(index string) *ResumeFollow

func NewResumeFollowFunc(tp elastictransport.Interface) NewResumeFollow {
	_ = "STUB: not implemented"
	return *new(NewResumeFollow)
}

func New(tp elastictransport.Interface) *ResumeFollow { _ = "STUB: not implemented"; return nil }

func (r *ResumeFollow) Raw(raw io.Reader) *ResumeFollow { _ = "STUB: not implemented"; return nil }

func (r *ResumeFollow) Request(req *Request) *ResumeFollow { _ = "STUB: not implemented"; return nil }

func (r *ResumeFollow) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResumeFollow) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResumeFollow) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ResumeFollow) Header(key, value string) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) _index(index string) *ResumeFollow { _ = "STUB: not implemented"; return nil }

func (r *ResumeFollow) MasterTimeout(duration string) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) ErrorTrace(errortrace bool) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) FilterPath(filterpaths ...string) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) Human(human bool) *ResumeFollow { _ = "STUB: not implemented"; return nil }

func (r *ResumeFollow) Pretty(pretty bool) *ResumeFollow { _ = "STUB: not implemented"; return nil }

func (r *ResumeFollow) MaxOutstandingReadRequests(maxoutstandingreadrequests int64) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) MaxOutstandingWriteRequests(maxoutstandingwriterequests int64) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) MaxReadRequestOperationCount(maxreadrequestoperationcount int64) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) MaxReadRequestSize(maxreadrequestsize string) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) MaxRetryDelay(duration types.DurationVariant) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) MaxWriteBufferCount(maxwritebuffercount int64) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) MaxWriteBufferSize(maxwritebuffersize string) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) MaxWriteRequestOperationCount(maxwriterequestoperationcount int64) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) MaxWriteRequestSize(maxwriterequestsize string) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResumeFollow) ReadPollTimeout(duration types.DurationVariant) *ResumeFollow {
	_ = "STUB: not implemented"
	return nil
}
