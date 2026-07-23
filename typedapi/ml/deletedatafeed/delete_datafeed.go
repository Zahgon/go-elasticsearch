package deletedatafeed

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

type DeleteDatafeed struct {
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

type NewDeleteDatafeed func(datafeedid string) *DeleteDatafeed

func NewDeleteDatafeedFunc(tp elastictransport.Interface) NewDeleteDatafeed {
	_ = "STUB: not implemented"
	return *new(NewDeleteDatafeed)
}

func New(tp elastictransport.Interface) *DeleteDatafeed { _ = "STUB: not implemented"; return nil }

func (r *DeleteDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDatafeed) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDatafeed) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDatafeed) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteDatafeed) Header(key, value string) *DeleteDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDatafeed) _datafeedid(datafeedid string) *DeleteDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDatafeed) Force(force bool) *DeleteDatafeed { _ = "STUB: not implemented"; return nil }

func (r *DeleteDatafeed) ErrorTrace(errortrace bool) *DeleteDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDatafeed) FilterPath(filterpaths ...string) *DeleteDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDatafeed) Human(human bool) *DeleteDatafeed { _ = "STUB: not implemented"; return nil }

func (r *DeleteDatafeed) Pretty(pretty bool) *DeleteDatafeed { _ = "STUB: not implemented"; return nil }
