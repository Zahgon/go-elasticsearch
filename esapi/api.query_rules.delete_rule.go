package esapi

import (
	"context"
	"net/http"
)

func newQueryRulesDeleteRuleFunc(t Transport) QueryRulesDeleteRule {
	_ = "STUB: not implemented"
	return *new(QueryRulesDeleteRule)
}

type QueryRulesDeleteRule func(rule_id string, ruleset_id string, o ...func(*QueryRulesDeleteRuleRequest)) (*Response, error)

type QueryRulesDeleteRuleRequest struct {
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

func (r QueryRulesDeleteRuleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f QueryRulesDeleteRule) WithContext(v context.Context) func(*QueryRulesDeleteRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRule) WithPretty() func(*QueryRulesDeleteRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRule) WithHuman() func(*QueryRulesDeleteRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRule) WithErrorTrace() func(*QueryRulesDeleteRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRule) WithFilterPath(v ...string) func(*QueryRulesDeleteRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRule) WithHeader(h map[string]string) func(*QueryRulesDeleteRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRule) WithOpaqueID(s string) func(*QueryRulesDeleteRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}
