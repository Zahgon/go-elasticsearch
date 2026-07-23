package listrulesets

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ListRulesets struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewListRulesets func() *ListRulesets

func NewListRulesetsFunc(tp elastictransport.Interface) NewListRulesets {
	_ = "STUB: not implemented"
	return *new(NewListRulesets)
}

func New(tp elastictransport.Interface) *ListRulesets { _ = "STUB: not implemented"; return nil }

func (r *ListRulesets) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListRulesets) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListRulesets) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ListRulesets) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ListRulesets) Header(key, value string) *ListRulesets {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListRulesets) From(from int) *ListRulesets { _ = "STUB: not implemented"; return nil }

func (r *ListRulesets) Size(size int) *ListRulesets { _ = "STUB: not implemented"; return nil }

func (r *ListRulesets) ErrorTrace(errortrace bool) *ListRulesets {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListRulesets) FilterPath(filterpaths ...string) *ListRulesets {
	_ = "STUB: not implemented"
	return nil
}

func (r *ListRulesets) Human(human bool) *ListRulesets { _ = "STUB: not implemented"; return nil }

func (r *ListRulesets) Pretty(pretty bool) *ListRulesets { _ = "STUB: not implemented"; return nil }
