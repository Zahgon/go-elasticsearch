package putruleset

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	rulesetidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutRuleset struct {
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

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutRuleset func(rulesetid string) *PutRuleset

func NewPutRulesetFunc(tp elastictransport.Interface) NewPutRuleset {
	_ = "STUB: not implemented"
	return *new(NewPutRuleset)
}

func New(tp elastictransport.Interface) *PutRuleset { _ = "STUB: not implemented"; return nil }

func (r *PutRuleset) Raw(raw io.Reader) *PutRuleset { _ = "STUB: not implemented"; return nil }

func (r *PutRuleset) Request(req *Request) *PutRuleset { _ = "STUB: not implemented"; return nil }

func (r *PutRuleset) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutRuleset) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutRuleset) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutRuleset) Header(key, value string) *PutRuleset { _ = "STUB: not implemented"; return nil }

func (r *PutRuleset) _rulesetid(rulesetid string) *PutRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRuleset) ErrorTrace(errortrace bool) *PutRuleset { _ = "STUB: not implemented"; return nil }

func (r *PutRuleset) FilterPath(filterpaths ...string) *PutRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutRuleset) Human(human bool) *PutRuleset { _ = "STUB: not implemented"; return nil }

func (r *PutRuleset) Pretty(pretty bool) *PutRuleset { _ = "STUB: not implemented"; return nil }

func (r *PutRuleset) Rules(rules ...types.QueryRuleVariant) *PutRuleset {
	_ = "STUB: not implemented"
	return nil
}
