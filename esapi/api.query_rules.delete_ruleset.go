package esapi

import (
	"context"
	"net/http"
)

func newQueryRulesDeleteRulesetFunc(t Transport) QueryRulesDeleteRuleset {
	_ = "STUB: not implemented"
	return *new(QueryRulesDeleteRuleset)
}

type QueryRulesDeleteRuleset func(ruleset_id string, o ...func(*QueryRulesDeleteRulesetRequest)) (*Response, error)

type QueryRulesDeleteRulesetRequest struct {
	RulesetID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r QueryRulesDeleteRulesetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f QueryRulesDeleteRuleset) WithContext(v context.Context) func(*QueryRulesDeleteRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRuleset) WithPretty() func(*QueryRulesDeleteRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRuleset) WithHuman() func(*QueryRulesDeleteRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRuleset) WithErrorTrace() func(*QueryRulesDeleteRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRuleset) WithFilterPath(v ...string) func(*QueryRulesDeleteRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRuleset) WithHeader(h map[string]string) func(*QueryRulesDeleteRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesDeleteRuleset) WithOpaqueID(s string) func(*QueryRulesDeleteRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}
