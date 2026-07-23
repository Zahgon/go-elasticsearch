package putautofollowpattern

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAutoFollowPattern struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutAutoFollowPattern func(name string) *PutAutoFollowPattern

func NewPutAutoFollowPatternFunc(tp elastictransport.Interface) NewPutAutoFollowPattern {
	_ = "STUB: not implemented"
	return *new(NewPutAutoFollowPattern)
}

func New(tp elastictransport.Interface) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) Raw(raw io.Reader) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) Request(req *Request) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAutoFollowPattern) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAutoFollowPattern) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAutoFollowPattern) Header(key, value string) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) _name(name string) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MasterTimeout(duration string) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) ErrorTrace(errortrace bool) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) FilterPath(filterpaths ...string) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) Human(human bool) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) Pretty(pretty bool) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) FollowIndexPattern(indexpattern string) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) LeaderIndexExclusionPatterns(indexpatterns ...string) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) LeaderIndexPatterns(indexpatterns ...string) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MaxOutstandingReadRequests(maxoutstandingreadrequests int) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MaxOutstandingWriteRequests(maxoutstandingwriterequests int) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MaxReadRequestOperationCount(maxreadrequestoperationcount int) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MaxReadRequestSize(bytesize types.ByteSizeVariant) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MaxRetryDelay(duration types.DurationVariant) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MaxWriteBufferCount(maxwritebuffercount int) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MaxWriteBufferSize(bytesize types.ByteSizeVariant) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MaxWriteRequestOperationCount(maxwriterequestoperationcount int) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) MaxWriteRequestSize(bytesize types.ByteSizeVariant) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) ReadPollTimeout(duration types.DurationVariant) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) RemoteCluster(remotecluster string) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) Settings(settings map[string]json.RawMessage) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoFollowPattern) AddSetting(key string, value json.RawMessage) *PutAutoFollowPattern {
	_ = "STUB: not implemented"
	return nil
}
