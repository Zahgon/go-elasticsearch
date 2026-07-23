package getsynonymrule

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	setidMask = iota + 1

	ruleidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetSynonymRule struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	setid  string
	ruleid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetSynonymRule func(setid, ruleid string) *GetSynonymRule

func NewGetSynonymRuleFunc(tp elastictransport.Interface) NewGetSynonymRule {
	_ = "STUB: not implemented"
	return *new(NewGetSynonymRule)
}

func New(tp elastictransport.Interface) *GetSynonymRule { _ = "STUB: not implemented"; return nil }

func (r *GetSynonymRule) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSynonymRule) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSynonymRule) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetSynonymRule) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetSynonymRule) Header(key, value string) *GetSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonymRule) _setid(setid string) *GetSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonymRule) _ruleid(ruleid string) *GetSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonymRule) ErrorTrace(errortrace bool) *GetSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonymRule) FilterPath(filterpaths ...string) *GetSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetSynonymRule) Human(human bool) *GetSynonymRule { _ = "STUB: not implemented"; return nil }

func (r *GetSynonymRule) Pretty(pretty bool) *GetSynonymRule { _ = "STUB: not implemented"; return nil }
