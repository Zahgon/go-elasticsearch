package addblock

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

type AddBlock struct {
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

type NewAddBlock func(index, block string) *AddBlock

func NewAddBlockFunc(tp elastictransport.Interface) NewAddBlock {
	_ = "STUB: not implemented"
	return *new(NewAddBlock)
}

func New(tp elastictransport.Interface) *AddBlock { _ = "STUB: not implemented"; return nil }

func (r *AddBlock) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AddBlock) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AddBlock) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r AddBlock) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *AddBlock) Header(key, value string) *AddBlock { _ = "STUB: not implemented"; return nil }

func (r *AddBlock) _index(index string) *AddBlock { _ = "STUB: not implemented"; return nil }

func (r *AddBlock) _block(block string) *AddBlock { _ = "STUB: not implemented"; return nil }

func (r *AddBlock) AllowNoIndices(allownoindices bool) *AddBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *AddBlock) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *AddBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *AddBlock) IgnoreUnavailable(ignoreunavailable bool) *AddBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *AddBlock) MasterTimeout(duration string) *AddBlock { _ = "STUB: not implemented"; return nil }

func (r *AddBlock) Timeout(duration string) *AddBlock { _ = "STUB: not implemented"; return nil }

func (r *AddBlock) ErrorTrace(errortrace bool) *AddBlock { _ = "STUB: not implemented"; return nil }

func (r *AddBlock) FilterPath(filterpaths ...string) *AddBlock {
	_ = "STUB: not implemented"
	return nil
}

func (r *AddBlock) Human(human bool) *AddBlock { _ = "STUB: not implemented"; return nil }

func (r *AddBlock) Pretty(pretty bool) *AddBlock { _ = "STUB: not implemented"; return nil }
