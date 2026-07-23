package getruleset

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	rulesetidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetRuleset struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	rulesetid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetRuleset func(rulesetid string) *GetRuleset

func NewGetRulesetFunc(tp elastictransport.Interface) NewGetRuleset {
	_ = "STUB: not implemented"
	return *new(NewGetRuleset)
}

func New(tp elastictransport.Interface) *GetRuleset { _ = "STUB: not implemented"; return nil }

func (r *GetRuleset) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRuleset) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRuleset) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRuleset) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetRuleset) Header(key, value string) *GetRuleset { _ = "STUB: not implemented"; return nil }

func (r *GetRuleset) _rulesetid(rulesetid string) *GetRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRuleset) ErrorTrace(errortrace bool) *GetRuleset { _ = "STUB: not implemented"; return nil }

func (r *GetRuleset) FilterPath(filterpaths ...string) *GetRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRuleset) Human(human bool) *GetRuleset { _ = "STUB: not implemented"; return nil }

func (r *GetRuleset) Pretty(pretty bool) *GetRuleset { _ = "STUB: not implemented"; return nil }
