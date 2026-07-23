package starttrainedmodeldeployment

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/deploymentallocationstate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/trainingpriority"
)

const (
	modelidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type StartTrainedModelDeployment struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	modelid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewStartTrainedModelDeployment func(modelid string) *StartTrainedModelDeployment

func NewStartTrainedModelDeploymentFunc(tp elastictransport.Interface) NewStartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return *new(NewStartTrainedModelDeployment)
}

func New(tp elastictransport.Interface) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) Raw(raw io.Reader) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) Request(req *Request) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartTrainedModelDeployment) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartTrainedModelDeployment) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *StartTrainedModelDeployment) Header(key, value string) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) _modelid(modelid string) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) CacheSize(bytesize string) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) DeploymentId(deploymentid string) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) NumberOfAllocations(numberofallocations int) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) Priority(priority trainingpriority.TrainingPriority) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) QueueCapacity(queuecapacity int) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) ThreadsPerAllocation(threadsperallocation int) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) Timeout(duration string) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) WaitFor(waitfor deploymentallocationstate.DeploymentAllocationState) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) ErrorTrace(errortrace bool) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) FilterPath(filterpaths ...string) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) Human(human bool) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) Pretty(pretty bool) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTrainedModelDeployment) AdaptiveAllocations(adaptiveallocations types.AdaptiveAllocationsSettingsVariant) *StartTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}
