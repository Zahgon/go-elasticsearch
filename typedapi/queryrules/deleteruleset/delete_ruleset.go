package deleteruleset

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

type DeleteRuleset struct {
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

type NewDeleteRuleset func(rulesetid string) *DeleteRuleset

func NewDeleteRulesetFunc(tp elastictransport.Interface) NewDeleteRuleset {
	_ = "STUB: not implemented"
	return *new(NewDeleteRuleset)
}

func New(tp elastictransport.Interface) *DeleteRuleset { _ = "STUB: not implemented"; return nil }

func (r *DeleteRuleset) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRuleset) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRuleset) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRuleset) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteRuleset) Header(key, value string) *DeleteRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRuleset) _rulesetid(rulesetid string) *DeleteRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRuleset) ErrorTrace(errortrace bool) *DeleteRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRuleset) FilterPath(filterpaths ...string) *DeleteRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRuleset) Human(human bool) *DeleteRuleset { _ = "STUB: not implemented"; return nil }

func (r *DeleteRuleset) Pretty(pretty bool) *DeleteRuleset { _ = "STUB: not implemented"; return nil }
