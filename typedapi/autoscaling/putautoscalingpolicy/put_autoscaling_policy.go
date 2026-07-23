package putautoscalingpolicy

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAutoscalingPolicy struct {
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

type NewPutAutoscalingPolicy func(name string) *PutAutoscalingPolicy

func NewPutAutoscalingPolicyFunc(tp elastictransport.Interface) NewPutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return *new(NewPutAutoscalingPolicy)
}

func New(tp elastictransport.Interface) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) Raw(raw io.Reader) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) Request(req *Request) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAutoscalingPolicy) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAutoscalingPolicy) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAutoscalingPolicy) Header(key, value string) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) _name(name string) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) MasterTimeout(duration string) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) Timeout(duration string) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) ErrorTrace(errortrace bool) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) FilterPath(filterpaths ...string) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) Human(human bool) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) Pretty(pretty bool) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) Deciders(deciders map[string]json.RawMessage) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) AddDecider(key string, value json.RawMessage) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAutoscalingPolicy) Roles(roles ...string) *PutAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}
