package putalibabacloud

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/alibabacloudservicetype"
)

const (
	tasktypeMask = iota + 1

	alibabacloudinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAlibabacloud struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                string
	alibabacloudinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutAlibabacloud func(tasktype, alibabacloudinferenceid string) *PutAlibabacloud

func NewPutAlibabacloudFunc(tp elastictransport.Interface) NewPutAlibabacloud {
	_ = "STUB: not implemented"
	return *new(NewPutAlibabacloud)
}

func New(tp elastictransport.Interface) *PutAlibabacloud { _ = "STUB: not implemented"; return nil }

func (r *PutAlibabacloud) Raw(raw io.Reader) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) Request(req *Request) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAlibabacloud) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAlibabacloud) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAlibabacloud) Header(key, value string) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) _tasktype(tasktype string) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) _alibabacloudinferenceid(alibabacloudinferenceid string) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) Timeout(duration string) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) ErrorTrace(errortrace bool) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) FilterPath(filterpaths ...string) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) Human(human bool) *PutAlibabacloud { _ = "STUB: not implemented"; return nil }

func (r *PutAlibabacloud) Pretty(pretty bool) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) Service(service alibabacloudservicetype.AlibabaCloudServiceType) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) ServiceSettings(servicesettings types.AlibabaCloudServiceSettingsVariant) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAlibabacloud) TaskSettings(tasksettings types.AlibabaCloudTaskSettingsVariant) *PutAlibabacloud {
	_ = "STUB: not implemented"
	return nil
}
