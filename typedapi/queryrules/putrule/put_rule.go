package putrule

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/queryruletype"
)

const (
	rulesetidMask = iota + 1

	ruleidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutRule struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	rulesetid string
	ruleid    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutRule func(rulesetid, ruleid string) *PutRule

func NewPutRuleFunc(tp elastictransport.Interface) NewPutRule {
	_ = "STUB: not implemented"
	return *new(NewPutRule)
}

func New(tp elastictransport.Interface) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) Raw(raw io.Reader) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) Request(req *Request) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutRule) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutRule) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutRule) Header(key, value string) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) _rulesetid(rulesetid string) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) _ruleid(ruleid string) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) ErrorTrace(errortrace bool) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) FilterPath(filterpaths ...string) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) Human(human bool) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) Pretty(pretty bool) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) Actions(actions types.QueryRuleActionsVariant) *PutRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRule) Criteria(criteria ...types.QueryRuleCriteriaVariant) *PutRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRule) Priority(priority int) *PutRule { _ = "STUB: not implemented"; return nil }

func (r *PutRule) Type(type_ queryruletype.QueryRuleType) *PutRule {
	_ = "STUB: not implemented"
	return nil
}
