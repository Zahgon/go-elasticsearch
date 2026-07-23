package getnodestats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetNodeStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetNodeStats func() *GetNodeStats

func NewGetNodeStatsFunc(tp elastictransport.Interface) NewGetNodeStats {
	_ = "STUB: not implemented"
	return *new(NewGetNodeStats)
}

func New(tp elastictransport.Interface) *GetNodeStats { _ = "STUB: not implemented"; return nil }

func (r *GetNodeStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetNodeStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetNodeStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetNodeStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetNodeStats) Header(key, value string) *GetNodeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetNodeStats) ErrorTrace(errortrace bool) *GetNodeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetNodeStats) FilterPath(filterpaths ...string) *GetNodeStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetNodeStats) Human(human bool) *GetNodeStats { _ = "STUB: not implemented"; return nil }

func (r *GetNodeStats) Pretty(pretty bool) *GetNodeStats { _ = "STUB: not implemented"; return nil }
