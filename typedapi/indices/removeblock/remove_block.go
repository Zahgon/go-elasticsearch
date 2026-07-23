package removeblock

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1

	blockMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RemoveBlock struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string
	block string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRemoveBlock func(index, block string) *RemoveBlock

func NewRemoveBlockFunc(tp elastictransport.Interface) NewRemoveBlock {
	_ = "STUB: not implemented"
	return *new(NewRemoveBlock)
}

func New(tp elastictransport.Interface) *RemoveBlock { _ = "STUB: not implemented"; return nil }

func (r *RemoveBlock) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RemoveBlock) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RemoveBlock) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RemoveBlock) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *RemoveBlock) Header(key, value string) *RemoveBlock { _ = "STUB: not implemented"; return nil }

func (r *RemoveBlock) _index(index string) *RemoveBlock { _ = "STUB: not implemented"; return nil }

func (r *RemoveBlock) _block(block string) *RemoveBlock { _ = "STUB: not implemented"; return nil }

func (r *RemoveBlock) AllowNoIndices(allownoindices bool) *RemoveBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoveBlock) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *RemoveBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoveBlock) IgnoreUnavailable(ignoreunavailable bool) *RemoveBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoveBlock) MasterTimeout(duration string) *RemoveBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoveBlock) Timeout(duration string) *RemoveBlock { _ = "STUB: not implemented"; return nil }

func (r *RemoveBlock) ErrorTrace(errortrace bool) *RemoveBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoveBlock) FilterPath(filterpaths ...string) *RemoveBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *RemoveBlock) Human(human bool) *RemoveBlock { _ = "STUB: not implemented"; return nil }

func (r *RemoveBlock) Pretty(pretty bool) *RemoveBlock { _ = "STUB: not implemented"; return nil }
