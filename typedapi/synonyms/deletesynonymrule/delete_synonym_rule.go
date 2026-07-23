package deletesynonymrule

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

type DeleteSynonymRule struct {
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

type NewDeleteSynonymRule func(setid, ruleid string) *DeleteSynonymRule

func NewDeleteSynonymRuleFunc(tp elastictransport.Interface) NewDeleteSynonymRule {
	_ = "STUB: not implemented"
	return *new(NewDeleteSynonymRule)
}

func New(tp elastictransport.Interface) *DeleteSynonymRule { _ = "STUB: not implemented"; return nil }

func (r *DeleteSynonymRule) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteSynonymRule) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteSynonymRule) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteSynonymRule) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteSynonymRule) Header(key, value string) *DeleteSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonymRule) _setid(setid string) *DeleteSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonymRule) _ruleid(ruleid string) *DeleteSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonymRule) Refresh(refresh bool) *DeleteSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonymRule) ErrorTrace(errortrace bool) *DeleteSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonymRule) FilterPath(filterpaths ...string) *DeleteSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonymRule) Human(human bool) *DeleteSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteSynonymRule) Pretty(pretty bool) *DeleteSynonymRule {
	_ = "STUB: not implemented"
	return nil
}
