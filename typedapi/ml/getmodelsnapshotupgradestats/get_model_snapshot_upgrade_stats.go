package getmodelsnapshotupgradestats

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

type GetModelSnapshotUpgradeStats struct {
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

type NewGetModelSnapshotUpgradeStats func(jobid, snapshotid string) *GetModelSnapshotUpgradeStats

func NewGetModelSnapshotUpgradeStatsFunc(tp elastictransport.Interface) NewGetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return *new(NewGetModelSnapshotUpgradeStats)
}

func New(tp elastictransport.Interface) *GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshotUpgradeStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetModelSnapshotUpgradeStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetModelSnapshotUpgradeStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetModelSnapshotUpgradeStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetModelSnapshotUpgradeStats) Header(key, value string) *GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshotUpgradeStats) _jobid(jobid string) *GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshotUpgradeStats) _snapshotid(snapshotid string) *GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshotUpgradeStats) AllowNoMatch(allownomatch bool) *GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshotUpgradeStats) ErrorTrace(errortrace bool) *GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshotUpgradeStats) FilterPath(filterpaths ...string) *GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshotUpgradeStats) Human(human bool) *GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetModelSnapshotUpgradeStats) Pretty(pretty bool) *GetModelSnapshotUpgradeStats {
	_ = "STUB: not implemented"
	return nil
}
