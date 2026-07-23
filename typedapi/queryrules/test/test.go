package test

import (
	gobytes "bytes"
	"context"
	"encoding/json"
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

type Test struct {
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

type NewTest func(rulesetid string) *Test

func NewTestFunc(tp elastictransport.Interface) NewTest {
	_ = "STUB: not implemented"
	return *new(NewTest)
}

func New(tp elastictransport.Interface) *Test { _ = "STUB: not implemented"; return nil }

func (r *Test) Raw(raw io.Reader) *Test { _ = "STUB: not implemented"; return nil }

func (r *Test) Request(req *Request) *Test { _ = "STUB: not implemented"; return nil }

func (r *Test) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Test) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Test) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Test) Header(key, value string) *Test { _ = "STUB: not implemented"; return nil }

func (r *Test) _rulesetid(rulesetid string) *Test { _ = "STUB: not implemented"; return nil }

func (r *Test) ErrorTrace(errortrace bool) *Test { _ = "STUB: not implemented"; return nil }

func (r *Test) FilterPath(filterpaths ...string) *Test { _ = "STUB: not implemented"; return nil }

func (r *Test) Human(human bool) *Test { _ = "STUB: not implemented"; return nil }

func (r *Test) Pretty(pretty bool) *Test { _ = "STUB: not implemented"; return nil }

func (r *Test) MatchCriteria(matchcriteria map[string]json.RawMessage) *Test {
	_ = "STUB: not implemented"
	return nil
}

func (r *Test) AddMatchCriterion(key string, value json.RawMessage) *Test {
	_ = "STUB: not implemented"
	return nil
}
