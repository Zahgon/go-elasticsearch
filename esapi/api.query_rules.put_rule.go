package esapi

import (
	"context"
	"io"
	"net/http"
)

func newQueryRulesPutRuleFunc(t Transport) QueryRulesPutRule {
	_ = "STUB: not implemented"
	return *new(QueryRulesPutRule)
}

type QueryRulesPutRule func(body io.Reader, rule_id string, ruleset_id string, o ...func(*QueryRulesPutRuleRequest)) (*Response, error)

type QueryRulesPutRuleRequest struct {
	Body io.Reader

	RuleID    string
	RulesetID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r QueryRulesPutRuleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f QueryRulesPutRule) WithContext(v context.Context) func(*QueryRulesPutRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRule) WithPretty() func(*QueryRulesPutRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRule) WithHuman() func(*QueryRulesPutRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRule) WithErrorTrace() func(*QueryRulesPutRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRule) WithFilterPath(v ...string) func(*QueryRulesPutRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRule) WithHeader(h map[string]string) func(*QueryRulesPutRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRule) WithOpaqueID(s string) func(*QueryRulesPutRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}
