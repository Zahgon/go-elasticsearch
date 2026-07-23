package follow

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

type Follow struct {
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

type NewFollow func(index string) *Follow

func NewFollowFunc(tp elastictransport.Interface) NewFollow {
	_ = "STUB: not implemented"
	return *new(NewFollow)
}

func New(tp elastictransport.Interface) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) Raw(raw io.Reader) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) Request(req *Request) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Follow) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Follow) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Follow) Header(key, value string) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) _index(index string) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) MasterTimeout(duration string) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) WaitForActiveShards(waitforactiveshards string) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) ErrorTrace(errortrace bool) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) FilterPath(filterpaths ...string) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) Human(human bool) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) Pretty(pretty bool) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) DataStreamName(datastreamname string) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) LeaderIndex(indexname string) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) MaxOutstandingReadRequests(maxoutstandingreadrequests int64) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) MaxOutstandingWriteRequests(maxoutstandingwriterequests int) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) MaxReadRequestOperationCount(maxreadrequestoperationcount int) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) MaxReadRequestSize(bytesize types.ByteSizeVariant) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) MaxRetryDelay(duration types.DurationVariant) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) MaxWriteBufferCount(maxwritebuffercount int) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) MaxWriteBufferSize(bytesize types.ByteSizeVariant) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) MaxWriteRequestOperationCount(maxwriterequestoperationcount int) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) MaxWriteRequestSize(bytesize types.ByteSizeVariant) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) ReadPollTimeout(duration types.DurationVariant) *Follow {
	_ = "STUB: not implemented"
	return nil
}

func (r *Follow) RemoteCluster(remotecluster string) *Follow { _ = "STUB: not implemented"; return nil }

func (r *Follow) Settings(settings types.IndexSettingsVariant) *Follow {
	_ = "STUB: not implemented"
	return nil
}
