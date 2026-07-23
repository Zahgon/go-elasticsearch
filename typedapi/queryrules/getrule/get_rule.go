package getrule

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

type GetRule struct {
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

type NewGetRule func(rulesetid, ruleid string) *GetRule

func NewGetRuleFunc(tp elastictransport.Interface) NewGetRule {
	_ = "STUB: not implemented"
	return *new(NewGetRule)
}

func New(tp elastictransport.Interface) *GetRule { _ = "STUB: not implemented"; return nil }

func (r *GetRule) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRule) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRule) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRule) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetRule) Header(key, value string) *GetRule { _ = "STUB: not implemented"; return nil }

func (r *GetRule) _rulesetid(rulesetid string) *GetRule { _ = "STUB: not implemented"; return nil }

func (r *GetRule) _ruleid(ruleid string) *GetRule { _ = "STUB: not implemented"; return nil }

func (r *GetRule) ErrorTrace(errortrace bool) *GetRule { _ = "STUB: not implemented"; return nil }

func (r *GetRule) FilterPath(filterpaths ...string) *GetRule { _ = "STUB: not implemented"; return nil }

func (r *GetRule) Human(human bool) *GetRule { _ = "STUB: not implemented"; return nil }

func (r *GetRule) Pretty(pretty bool) *GetRule { _ = "STUB: not implemented"; return nil }
