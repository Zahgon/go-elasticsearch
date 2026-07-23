package importdanglingindex

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

type ImportDanglingIndex struct {
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

type NewImportDanglingIndex func(indexuuid string) *ImportDanglingIndex

func NewImportDanglingIndexFunc(tp elastictransport.Interface) NewImportDanglingIndex {
	_ = "STUB: not implemented"
	return *new(NewImportDanglingIndex)
}

func New(tp elastictransport.Interface) *ImportDanglingIndex { _ = "STUB: not implemented"; return nil }

func (r *ImportDanglingIndex) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ImportDanglingIndex) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ImportDanglingIndex) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ImportDanglingIndex) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ImportDanglingIndex) Header(key, value string) *ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ImportDanglingIndex) _indexuuid(indexuuid string) *ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ImportDanglingIndex) AcceptDataLoss(acceptdataloss bool) *ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ImportDanglingIndex) MasterTimeout(duration string) *ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ImportDanglingIndex) Timeout(duration string) *ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ImportDanglingIndex) ErrorTrace(errortrace bool) *ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ImportDanglingIndex) FilterPath(filterpaths ...string) *ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ImportDanglingIndex) Human(human bool) *ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}

func (r *ImportDanglingIndex) Pretty(pretty bool) *ImportDanglingIndex {
	_ = "STUB: not implemented"
	return nil
}
