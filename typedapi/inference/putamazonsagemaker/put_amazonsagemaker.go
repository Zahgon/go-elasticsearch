package putamazonsagemaker

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/amazonsagemakerservicetype"
)

const (
	tasktypeMask = iota + 1

	amazonsagemakerinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAmazonsagemaker struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                   string
	amazonsagemakerinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutAmazonsagemaker func(tasktype, amazonsagemakerinferenceid string) *PutAmazonsagemaker

func NewPutAmazonsagemakerFunc(tp elastictransport.Interface) NewPutAmazonsagemaker {
	_ = "STUB: not implemented"
	return *new(NewPutAmazonsagemaker)
}

func New(tp elastictransport.Interface) *PutAmazonsagemaker { _ = "STUB: not implemented"; return nil }

func (r *PutAmazonsagemaker) Raw(raw io.Reader) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) Request(req *Request) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAmazonsagemaker) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutAmazonsagemaker) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutAmazonsagemaker) Header(key, value string) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) _tasktype(tasktype string) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) _amazonsagemakerinferenceid(amazonsagemakerinferenceid string) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) Timeout(duration string) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) ErrorTrace(errortrace bool) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) FilterPath(filterpaths ...string) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) Human(human bool) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) Pretty(pretty bool) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) Service(service amazonsagemakerservicetype.AmazonSageMakerServiceType) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) ServiceSettings(servicesettings types.AmazonSageMakerServiceSettingsVariant) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutAmazonsagemaker) TaskSettings(tasksettings types.AmazonSageMakerTaskSettingsVariant) *PutAmazonsagemaker {
	_ = "STUB: not implemented"
	return nil
}
