package list

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/groupby"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type List struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewList func() *List

func NewListFunc(tp elastictransport.Interface) NewList {
	_ = "STUB: not implemented"
	return *new(NewList)
}

func New(tp elastictransport.Interface) *List { _ = "STUB: not implemented"; return nil }

func (r *List) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r List) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r List) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r List) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *List) Header(key, value string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Actions(actions ...string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Detailed(detailed bool) *List { _ = "STUB: not implemented"; return nil }

func (r *List) GroupBy(groupby groupby.GroupBy) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Nodes(nodeids ...string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) ParentTaskId(id string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Timeout(duration string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) WaitForCompletion(waitforcompletion bool) *List {
	_ = "STUB: not implemented"
	return nil
}

func (r *List) ErrorTrace(errortrace bool) *List { _ = "STUB: not implemented"; return nil }

func (r *List) FilterPath(filterpaths ...string) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Human(human bool) *List { _ = "STUB: not implemented"; return nil }

func (r *List) Pretty(pretty bool) *List { _ = "STUB: not implemented"; return nil }
