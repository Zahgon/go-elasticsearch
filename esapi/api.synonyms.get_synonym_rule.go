package esapi

import (
	"context"
	"net/http"
)

func newSynonymsGetSynonymRuleFunc(t Transport) SynonymsGetSynonymRule {
	_ = "STUB: not implemented"
	return *new(SynonymsGetSynonymRule)
}

type SynonymsGetSynonymRule func(rule_id string, set_id string, o ...func(*SynonymsGetSynonymRuleRequest)) (*Response, error)

type SynonymsGetSynonymRuleRequest struct {
	RuleID string
	SetID  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SynonymsGetSynonymRuleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SynonymsGetSynonymRule) WithContext(v context.Context) func(*SynonymsGetSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymRule) WithPretty() func(*SynonymsGetSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymRule) WithHuman() func(*SynonymsGetSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymRule) WithErrorTrace() func(*SynonymsGetSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymRule) WithFilterPath(v ...string) func(*SynonymsGetSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymRule) WithHeader(h map[string]string) func(*SynonymsGetSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsGetSynonymRule) WithOpaqueID(s string) func(*SynonymsGetSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}
