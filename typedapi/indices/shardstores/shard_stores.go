package shardstores

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shardstorestatus"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ShardStores struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewShardStores func() *ShardStores

func NewShardStoresFunc(tp elastictransport.Interface) NewShardStores {
	_ = "STUB: not implemented"
	return *new(NewShardStores)
}

func New(tp elastictransport.Interface) *ShardStores { _ = "STUB: not implemented"; return nil }

func (r *ShardStores) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ShardStores) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ShardStores) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ShardStores) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ShardStores) Header(key, value string) *ShardStores { _ = "STUB: not implemented"; return nil }

func (r *ShardStores) Index(index string) *ShardStores { _ = "STUB: not implemented"; return nil }

func (r *ShardStores) AllowNoIndices(allownoindices bool) *ShardStores {
	_ = "STUB: not implemented"
	return nil
}

func (r *ShardStores) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ShardStores {
	_ = "STUB: not implemented"
	return nil
}

func (r *ShardStores) IgnoreUnavailable(ignoreunavailable bool) *ShardStores {
	_ = "STUB: not implemented"
	return nil
}

func (r *ShardStores) Status(statuses ...shardstorestatus.ShardStoreStatus) *ShardStores {
	_ = "STUB: not implemented"
	return nil
}

func (r *ShardStores) ErrorTrace(errortrace bool) *ShardStores {
	_ = "STUB: not implemented"
	return nil
}

func (r *ShardStores) FilterPath(filterpaths ...string) *ShardStores {
	_ = "STUB: not implemented"
	return nil
}

func (r *ShardStores) Human(human bool) *ShardStores { _ = "STUB: not implemented"; return nil }

func (r *ShardStores) Pretty(pretty bool) *ShardStores { _ = "STUB: not implemented"; return nil }
