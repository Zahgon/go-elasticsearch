package putamazonbedrock

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonbedrockservicetype"
)

const (
	tasktypeMask = iota + 1

	amazonbedrockinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAmazonbedrock struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                 string
	amazonbedrockinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutAmazonbedrock func(tasktype, amazonbedrockinferenceid string) *PutAmazonbedrock

func NewPutAmazonbedrockFunc(tp elastictransport.Interface) NewPutAmazonbedrock {
	_ = "STUB: not implemented"
	return *new(NewPutAmazonbedrock)
}

func New(tp elastictransport.Interface) *PutAmazonbedrock { _ = "STUB: not implemented"; return nil }

func (r *PutAmazonbedrock) Raw(raw io.Reader) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) Request(req *Request) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAmazonbedrock) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAmazonbedrock) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAmazonbedrock) Header(key, value string) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) _tasktype(tasktype string) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) _amazonbedrockinferenceid(amazonbedrockinferenceid string) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) Timeout(duration string) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) ErrorTrace(errortrace bool) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) FilterPath(filterpaths ...string) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) Human(human bool) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) Pretty(pretty bool) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) Service(service amazonbedrockservicetype.AmazonBedrockServiceType) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) ServiceSettings(servicesettings types.AmazonBedrockServiceSettingsVariant) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonbedrock) TaskSettings(tasksettings types.AmazonBedrockTaskSettingsVariant) *PutAmazonbedrock {
	_ = "STUB: not implemented"
	return nil
}
