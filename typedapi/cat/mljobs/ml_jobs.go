package mljobs

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/catanomalydetectorcolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	jobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MlJobs struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewMlJobs func() *MlJobs

func NewMlJobsFunc(tp elastictransport.Interface) NewMlJobs {
	_ = "STUB: not implemented"
	return *new(NewMlJobs)
}

func New(tp elastictransport.Interface) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MlJobs) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MlJobs) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r MlJobs) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *MlJobs) Header(key, value string) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) JobId(jobid string) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) AllowNoMatch(allownomatch bool) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) H(catanomalydetectorcolumns ...catanomalydetectorcolumn.CatAnomalyDetectorColumn) *MlJobs {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlJobs) S(catanomalydetectorcolumns ...catanomalydetectorcolumn.CatAnomalyDetectorColumn) *MlJobs {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlJobs) Bytes(bytes bytes.Bytes) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) Format(format string) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) Help(help bool) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) Time(time timeunit.TimeUnit) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) V(v bool) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) ErrorTrace(errortrace bool) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) FilterPath(filterpaths ...string) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) Human(human bool) *MlJobs { _ = "STUB: not implemented"; return nil }

func (r *MlJobs) Pretty(pretty bool) *MlJobs { _ = "STUB: not implemented"; return nil }
