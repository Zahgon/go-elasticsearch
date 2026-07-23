package getmodelsnapshots

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	jobidMask = iota + 1

	snapshotidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetModelSnapshots struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid      string
	snapshotid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetModelSnapshots func(jobid string) *GetModelSnapshots

func NewGetModelSnapshotsFunc(tp elastictransport.Interface) NewGetModelSnapshots {
	_ = "STUB: not implemented"
	return *new(NewGetModelSnapshots)
}

func New(tp elastictransport.Interface) *GetModelSnapshots { _ = "STUB: not implemented"; return nil }

func (r *GetModelSnapshots) Raw(raw io.Reader) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) Request(req *Request) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetModelSnapshots) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetModelSnapshots) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *GetModelSnapshots) Header(key, value string) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) _jobid(jobid string) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) SnapshotId(snapshotid string) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) From(from int) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) Size(size int) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) ErrorTrace(errortrace bool) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) FilterPath(filterpaths ...string) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) Human(human bool) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) Pretty(pretty bool) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) Desc(desc bool) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) End(datetime types.DateTimeVariant) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) Page(page types.PageVariant) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) Sort(field string) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshots) Start(datetime types.DateTimeVariant) *GetModelSnapshots {
	_ = "STUB: not implemented"
	return nil
}
