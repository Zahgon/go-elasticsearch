package putsynonymrule

import (
	gobytes "bytes"
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

type PutSynonymRule struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	setid  string
	ruleid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutSynonymRule func(setid, ruleid string) *PutSynonymRule

func NewPutSynonymRuleFunc(tp elastictransport.Interface) NewPutSynonymRule {
	_ = "STUB: not implemented"
	return *new(NewPutSynonymRule)
}

func New(tp elastictransport.Interface) *PutSynonymRule { _ = "STUB: not implemented"; return nil }

func (r *PutSynonymRule) Raw(raw io.Reader) *PutSynonymRule { _ = "STUB: not implemented"; return nil }

func (r *PutSynonymRule) Request(req *Request) *PutSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSynonymRule) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutSynonymRule) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutSynonymRule) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutSynonymRule) Header(key, value string) *PutSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSynonymRule) _setid(setid string) *PutSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSynonymRule) _ruleid(ruleid string) *PutSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSynonymRule) Refresh(refresh bool) *PutSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSynonymRule) ErrorTrace(errortrace bool) *PutSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSynonymRule) FilterPath(filterpaths ...string) *PutSynonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutSynonymRule) Human(human bool) *PutSynonymRule { _ = "STUB: not implemented"; return nil }

func (r *PutSynonymRule) Pretty(pretty bool) *PutSynonymRule { _ = "STUB: not implemented"; return nil }

func (r *PutSynonymRule) Synonyms(synonymstring string) *PutSynonymRule {
	_ = "STUB: not implemented"
	return nil
}
