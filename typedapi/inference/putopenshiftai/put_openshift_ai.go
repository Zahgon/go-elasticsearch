package putopenshiftai

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openshiftaiservicetype"
)

const (
	tasktypeMask = iota + 1

	openshiftaiinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutOpenshiftAi struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype               string
	openshiftaiinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutOpenshiftAi func(tasktype, openshiftaiinferenceid string) *PutOpenshiftAi

func NewPutOpenshiftAiFunc(tp elastictransport.Interface) NewPutOpenshiftAi {
	_ = "STUB: not implemented"
	return *new(NewPutOpenshiftAi)
}

func New(tp elastictransport.Interface) *PutOpenshiftAi { _ = "STUB: not implemented"; return nil }

func (r *PutOpenshiftAi) Raw(raw io.Reader) *PutOpenshiftAi { _ = "STUB: not implemented"; return nil }

func (r *PutOpenshiftAi) Request(req *Request) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutOpenshiftAi) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutOpenshiftAi) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutOpenshiftAi) Header(key, value string) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) _tasktype(tasktype string) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) _openshiftaiinferenceid(openshiftaiinferenceid string) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) Timeout(duration string) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) ErrorTrace(errortrace bool) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) FilterPath(filterpaths ...string) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) Human(human bool) *PutOpenshiftAi { _ = "STUB: not implemented"; return nil }

func (r *PutOpenshiftAi) Pretty(pretty bool) *PutOpenshiftAi { _ = "STUB: not implemented"; return nil }

func (r *PutOpenshiftAi) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) Service(service openshiftaiservicetype.OpenShiftAiServiceType) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) ServiceSettings(servicesettings types.OpenShiftAiServiceSettingsVariant) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutOpenshiftAi) TaskSettings(tasksettings types.OpenShiftAiTaskSettingsVariant) *PutOpenshiftAi {
	_ = "STUB: not implemented"
	return nil
}
