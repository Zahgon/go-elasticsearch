package getdatafeedstats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	datafeedidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetDatafeedStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	datafeedid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetDatafeedStats func() *GetDatafeedStats

func NewGetDatafeedStatsFunc(tp elastictransport.Interface) NewGetDatafeedStats {
	_ = "STUB: not implemented"
	return *new(NewGetDatafeedStats)
}

func New(tp elastictransport.Interface) *GetDatafeedStats { _ = "STUB: not implemented"; return nil }

func (r *GetDatafeedStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDatafeedStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDatafeedStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDatafeedStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDatafeedStats) Header(key, value string) *GetDatafeedStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeedStats) DatafeedId(datafeedid string) *GetDatafeedStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeedStats) AllowNoMatch(allownomatch bool) *GetDatafeedStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeedStats) ErrorTrace(errortrace bool) *GetDatafeedStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeedStats) FilterPath(filterpaths ...string) *GetDatafeedStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeedStats) Human(human bool) *GetDatafeedStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeedStats) Pretty(pretty bool) *GetDatafeedStats {
	_ = "STUB: not implemented"
	return nil
}
