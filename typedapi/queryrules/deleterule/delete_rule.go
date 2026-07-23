package deleterule

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

	ruleidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteRule struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	rulesetid string
	ruleid    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteRule func(rulesetid, ruleid string) *DeleteRule

func NewDeleteRuleFunc(tp elastictransport.Interface) NewDeleteRule {
	_ = "STUB: not implemented"
	return *new(NewDeleteRule)
}

func New(tp elastictransport.Interface) *DeleteRule { _ = "STUB: not implemented"; return nil }

func (r *DeleteRule) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRule) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRule) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteRule) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteRule) Header(key, value string) *DeleteRule { _ = "STUB: not implemented"; return nil }

func (r *DeleteRule) _rulesetid(rulesetid string) *DeleteRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRule) _ruleid(ruleid string) *DeleteRule { _ = "STUB: not implemented"; return nil }

func (r *DeleteRule) ErrorTrace(errortrace bool) *DeleteRule { _ = "STUB: not implemented"; return nil }

func (r *DeleteRule) FilterPath(filterpaths ...string) *DeleteRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteRule) Human(human bool) *DeleteRule { _ = "STUB: not implemented"; return nil }

func (r *DeleteRule) Pretty(pretty bool) *DeleteRule { _ = "STUB: not implemented"; return nil }
