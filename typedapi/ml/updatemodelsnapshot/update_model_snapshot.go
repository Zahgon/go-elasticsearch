package updatemodelsnapshot

import (
	gobytes "bytes"
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

type UpdateModelSnapshot struct {
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

type NewUpdateModelSnapshot func(jobid, snapshotid string) *UpdateModelSnapshot

func NewUpdateModelSnapshotFunc(tp elastictransport.Interface) NewUpdateModelSnapshot {
	_ = "STUB: not implemented"
	return *new(NewUpdateModelSnapshot)
}

func New(tp elastictransport.Interface) *UpdateModelSnapshot { _ = "STUB: not implemented"; return nil }

func (r *UpdateModelSnapshot) Raw(raw io.Reader) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) Request(req *Request) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateModelSnapshot) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateModelSnapshot) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateModelSnapshot) Header(key, value string) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) _jobid(jobid string) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) _snapshotid(snapshotid string) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) ErrorTrace(errortrace bool) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) FilterPath(filterpaths ...string) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) Human(human bool) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) Pretty(pretty bool) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) Description(description string) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateModelSnapshot) Retain(retain bool) *UpdateModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}
