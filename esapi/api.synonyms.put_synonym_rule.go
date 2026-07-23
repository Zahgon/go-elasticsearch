package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSynonymsPutSynonymRuleFunc(t Transport) SynonymsPutSynonymRule {
	_ = "STUB: not implemented"
	return *new(SynonymsPutSynonymRule)
}

type SynonymsPutSynonymRule func(body io.Reader, rule_id string, set_id string, o ...func(*SynonymsPutSynonymRuleRequest)) (*Response, error)

type SynonymsPutSynonymRuleRequest struct {
	Body io.Reader

	RuleID string
	SetID  string

	Refresh *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SynonymsPutSynonymRuleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SynonymsPutSynonymRule) WithContext(v context.Context) func(*SynonymsPutSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonymRule) WithRefresh(v bool) func(*SynonymsPutSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonymRule) WithPretty() func(*SynonymsPutSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonymRule) WithHuman() func(*SynonymsPutSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonymRule) WithErrorTrace() func(*SynonymsPutSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonymRule) WithFilterPath(v ...string) func(*SynonymsPutSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonymRule) WithHeader(h map[string]string) func(*SynonymsPutSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsPutSynonymRule) WithOpaqueID(s string) func(*SynonymsPutSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}
