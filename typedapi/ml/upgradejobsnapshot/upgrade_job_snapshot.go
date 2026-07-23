package upgradejobsnapshot

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

type UpgradeJobSnapshot struct {
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

type NewUpgradeJobSnapshot func(jobid, snapshotid string) *UpgradeJobSnapshot

func NewUpgradeJobSnapshotFunc(tp elastictransport.Interface) NewUpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return *new(NewUpgradeJobSnapshot)
}

func New(tp elastictransport.Interface) *UpgradeJobSnapshot { _ = "STUB: not implemented"; return nil }

func (r *UpgradeJobSnapshot) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpgradeJobSnapshot) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpgradeJobSnapshot) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpgradeJobSnapshot) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *UpgradeJobSnapshot) Header(key, value string) *UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeJobSnapshot) _jobid(jobid string) *UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeJobSnapshot) _snapshotid(snapshotid string) *UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeJobSnapshot) WaitForCompletion(waitforcompletion bool) *UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeJobSnapshot) Timeout(duration string) *UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeJobSnapshot) ErrorTrace(errortrace bool) *UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeJobSnapshot) FilterPath(filterpaths ...string) *UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeJobSnapshot) Human(human bool) *UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpgradeJobSnapshot) Pretty(pretty bool) *UpgradeJobSnapshot {
	_ = "STUB: not implemented"
	return nil
}
