package esapi

import (
	"context"
	"io"
	"net/http"
)

func newQueryRulesPutRulesetFunc(t Transport) QueryRulesPutRuleset {
	_ = "STUB: not implemented"
	return *new(QueryRulesPutRuleset)
}

type QueryRulesPutRuleset func(body io.Reader, ruleset_id string, o ...func(*QueryRulesPutRulesetRequest)) (*Response, error)

type QueryRulesPutRulesetRequest struct {
	Body io.Reader

	RulesetID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r QueryRulesPutRulesetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f QueryRulesPutRuleset) WithContext(v context.Context) func(*QueryRulesPutRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRuleset) WithPretty() func(*QueryRulesPutRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRuleset) WithHuman() func(*QueryRulesPutRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRuleset) WithErrorTrace() func(*QueryRulesPutRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRuleset) WithFilterPath(v ...string) func(*QueryRulesPutRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRuleset) WithHeader(h map[string]string) func(*QueryRulesPutRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesPutRuleset) WithOpaqueID(s string) func(*QueryRulesPutRulesetRequest) {
	_ = "STUB: not implemented"
	return nil
}
