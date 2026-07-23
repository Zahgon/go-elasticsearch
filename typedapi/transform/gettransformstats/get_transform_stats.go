package gettransformstats

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	transformidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetTransformStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	transformid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetTransformStats func(transformid string) *GetTransformStats

func NewGetTransformStatsFunc(tp elastictransport.Interface) NewGetTransformStats {
	_ = "STUB: not implemented"
	return *new(NewGetTransformStats)
}

func New(tp elastictransport.Interface) *GetTransformStats { _ = "STUB: not implemented"; return nil }

func (r *GetTransformStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTransformStats) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTransformStats) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTransformStats) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetTransformStats) Header(key, value string) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransformStats) _transformid(transformid string) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransformStats) AllowNoMatch(allownomatch bool) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransformStats) From(from string) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransformStats) Size(size string) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransformStats) Timeout(duration string) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransformStats) ErrorTrace(errortrace bool) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransformStats) FilterPath(filterpaths ...string) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransformStats) Human(human bool) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransformStats) Pretty(pretty bool) *GetTransformStats {
	_ = "STUB: not implemented"
	return nil
}
