package esapi

import (
	"context"
	"net/http"
)

func newQueryRulesGetRulesetFunc(t Transport) QueryRulesGetRuleset {
	_ = "STUB: not implemented"
	return *new(QueryRulesGetRuleset)
}

type QueryRulesGetRuleset func(ruleset_id string, o ...func(*QueryRulesGetRulesetRequest)) (*Response, error)

type QueryRulesGetRulesetRequest struct {
	RulesetID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r QueryRulesGetRulesetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f QueryRulesGetRuleset) WithContext(v context.Context) func(*QueryRulesGetRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRuleset) WithPretty() func(*QueryRulesGetRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRuleset) WithHuman() func(*QueryRulesGetRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRuleset) WithErrorTrace() func(*QueryRulesGetRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRuleset) WithFilterPath(v ...string) func(*QueryRulesGetRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRuleset) WithHeader(h map[string]string) func(*QueryRulesGetRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesGetRuleset) WithOpaqueID(s string) func(*QueryRulesGetRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}
