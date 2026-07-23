package putgoogleaistudio

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/googleaiservicetype"
)

const (
	tasktypeMask = iota + 1

	googleaistudioinferenceidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutGoogleaistudio struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	tasktype                  string
	googleaistudioinferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutGoogleaistudio func(tasktype, googleaistudioinferenceid string) *PutGoogleaistudio

func NewPutGoogleaistudioFunc(tp elastictransport.Interface) NewPutGoogleaistudio {
	_ = "STUB: not implemented"
	return *new(NewPutGoogleaistudio)
}

func New(tp elastictransport.Interface) *PutGoogleaistudio { _ = "STUB: not implemented"; return nil }

func (r *PutGoogleaistudio) Raw(raw io.Reader) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) Request(req *Request) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutGoogleaistudio) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutGoogleaistudio) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutGoogleaistudio) Header(key, value string) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) _tasktype(tasktype string) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) _googleaistudioinferenceid(googleaistudioinferenceid string) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) Timeout(duration string) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) ErrorTrace(errortrace bool) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) FilterPath(filterpaths ...string) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) Human(human bool) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) Pretty(pretty bool) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) Service(service googleaiservicetype.GoogleAiServiceType) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutGoogleaistudio) ServiceSettings(servicesettings types.GoogleAiStudioServiceSettingsVariant) *PutGoogleaistudio {
	_ = "STUB: not implemented"
	return nil
}
