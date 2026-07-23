package esapi

import (
	"context"
	"net/http"
)

func newQueryRulesGetRuleFunc(t Transport) QueryRulesGetRule {
	_ = "STUB: not implemented"
	return *new(QueryRulesGetRule)
}

type QueryRulesGetRule func(rule_id string, ruleset_id string, o ...func(*QueryRulesGetRuleRequest)) (*Response, error)

type QueryRulesGetRuleRequest struct {
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

func (r QueryRulesGetRuleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f QueryRulesGetRule) WithContext(v context.Context) func(*QueryRulesGetRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRule) WithPretty() func(*QueryRulesGetRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRule) WithHuman() func(*QueryRulesGetRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRule) WithErrorTrace() func(*QueryRulesGetRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRule) WithFilterPath(v ...string) func(*QueryRulesGetRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRule) WithHeader(h map[string]string) func(*QueryRulesGetRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRule) WithOpaqueID(s string) func(*QueryRulesGetRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}
