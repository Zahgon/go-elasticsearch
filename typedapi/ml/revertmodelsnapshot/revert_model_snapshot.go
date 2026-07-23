package revertmodelsnapshot

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

type RevertModelSnapshot struct {
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

type NewRevertModelSnapshot func(jobid, snapshotid string) *RevertModelSnapshot

func NewRevertModelSnapshotFunc(tp elastictransport.Interface) NewRevertModelSnapshot {
	_ = "STUB: not implemented"
	return *new(NewRevertModelSnapshot)
}

func New(tp elastictransport.Interface) *RevertModelSnapshot { _ = "STUB: not implemented"; return nil }

func (r *RevertModelSnapshot) Raw(raw io.Reader) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *RevertModelSnapshot) Request(req *Request) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *RevertModelSnapshot) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RevertModelSnapshot) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RevertModelSnapshot) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RevertModelSnapshot) Header(key, value string) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *RevertModelSnapshot) _jobid(jobid string) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *RevertModelSnapshot) _snapshotid(snapshotid string) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *RevertModelSnapshot) ErrorTrace(errortrace bool) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *RevertModelSnapshot) FilterPath(filterpaths ...string) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *RevertModelSnapshot) Human(human bool) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *RevertModelSnapshot) Pretty(pretty bool) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *RevertModelSnapshot) DeleteInterveningResults(deleteinterveningresults bool) *RevertModelSnapshot {
	_ = "STUB: not implemented"
	return nil
}
