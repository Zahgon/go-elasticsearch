package deletedanglingindex

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexuuidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteDanglingIndex struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	indexuuid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteDanglingIndex func(indexuuid string) *DeleteDanglingIndex

func NewDeleteDanglingIndexFunc(tp elastictransport.Interface) NewDeleteDanglingIndex {
	_ = "STUB: not implemented"
	return *new(NewDeleteDanglingIndex)
}

func New(tp elastictransport.Interface) *DeleteDanglingIndex { _ = "STUB: not implemented"; return nil }

func (r *DeleteDanglingIndex) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDanglingIndex) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDanglingIndex) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteDanglingIndex) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteDanglingIndex) Header(key, value string) *DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDanglingIndex) _indexuuid(indexuuid string) *DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDanglingIndex) AcceptDataLoss(acceptdataloss bool) *DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDanglingIndex) MasterTimeout(duration string) *DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDanglingIndex) Timeout(duration string) *DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDanglingIndex) ErrorTrace(errortrace bool) *DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDanglingIndex) FilterPath(filterpaths ...string) *DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDanglingIndex) Human(human bool) *DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteDanglingIndex) Pretty(pretty bool) *DeleteDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}
