package getdatafeeds

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

type GetDatafeeds struct {
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

type NewGetDatafeeds func() *GetDatafeeds

func NewGetDatafeedsFunc(tp elastictransport.Interface) NewGetDatafeeds {
	_ = "STUB: not implemented"
	return *new(NewGetDatafeeds)
}

func New(tp elastictransport.Interface) *GetDatafeeds { _ = "STUB: not implemented"; return nil }

func (r *GetDatafeeds) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDatafeeds) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDatafeeds) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetDatafeeds) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetDatafeeds) Header(key, value string) *GetDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeeds) DatafeedId(datafeedid string) *GetDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeeds) AllowNoMatch(allownomatch bool) *GetDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeeds) ExcludeGenerated(excludegenerated bool) *GetDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeeds) ErrorTrace(errortrace bool) *GetDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeeds) FilterPath(filterpaths ...string) *GetDatafeeds {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetDatafeeds) Human(human bool) *GetDatafeeds { _ = "STUB: not implemented"; return nil }

func (r *GetDatafeeds) Pretty(pretty bool) *GetDatafeeds { _ = "STUB: not implemented"; return nil }
