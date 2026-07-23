package deletemodelsnapshot

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	jobidMask = iota + 1

	snapshotidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteModelSnapshot struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	jobid      string
	snapshotid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteModelSnapshot func(jobid, snapshotid string) *DeleteModelSnapshot

func NewDeleteModelSnapshotFunc(tp elastictransport.Interface) NewDeleteModelSnapshot {
	_ = "STUB: not implemented"
	return *new(NewDeleteModelSnapshot)
}

func New(tp elastictransport.Interface) *DeleteModelSnapshot { _ = "STUB: not implemented"; return nil }

func (r *DeleteModelSnapshot) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteModelSnapshot) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteModelSnapshot) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteModelSnapshot) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteModelSnapshot) Header(key, value string) *DeleteModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteModelSnapshot) _jobid(jobid string) *DeleteModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteModelSnapshot) _snapshotid(snapshotid string) *DeleteModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteModelSnapshot) ErrorTrace(errortrace bool) *DeleteModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteModelSnapshot) FilterPath(filterpaths ...string) *DeleteModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteModelSnapshot) Human(human bool) *DeleteModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteModelSnapshot) Pretty(pretty bool) *DeleteModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}
