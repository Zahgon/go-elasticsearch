package gettrainedmodelsstats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	modelidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetTrainedModelsStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	modelid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetTrainedModelsStats func() *GetTrainedModelsStats

func NewGetTrainedModelsStatsFunc(tp elastictransport.Interface) NewGetTrainedModelsStats {
	_ = "STUB: not implemented"
	return *new(NewGetTrainedModelsStats)
}

func New(tp elastictransport.Interface) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModelsStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTrainedModelsStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTrainedModelsStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTrainedModelsStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetTrainedModelsStats) Header(key, value string) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModelsStats) ModelId(modelid string) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModelsStats) AllowNoMatch(allownomatch bool) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModelsStats) From(from int) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModelsStats) Size(size int) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModelsStats) ErrorTrace(errortrace bool) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModelsStats) FilterPath(filterpaths ...string) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModelsStats) Human(human bool) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModelsStats) Pretty(pretty bool) *GetTrainedModelsStats {
	_ = "STUB: not implemented"
	return nil
}
