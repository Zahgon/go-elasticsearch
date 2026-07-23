package esapi

import (
	"context"
	"net/http"
)

func newSynonymsDeleteSynonymRuleFunc(t Transport) SynonymsDeleteSynonymRule {
	_ = "STUB: not implemented"
	return *new(SynonymsDeleteSynonymRule)
}

type SynonymsDeleteSynonymRule func(rule_id string, set_id string, o ...func(*SynonymsDeleteSynonymRuleRequest)) (*Response, error)

type SynonymsDeleteSynonymRuleRequest struct {
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

func (r SynonymsDeleteSynonymRuleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SynonymsDeleteSynonymRule) WithContext(v context.Context) func(*SynonymsDeleteSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonymRule) WithRefresh(v bool) func(*SynonymsDeleteSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonymRule) WithPretty() func(*SynonymsDeleteSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonymRule) WithHuman() func(*SynonymsDeleteSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonymRule) WithErrorTrace() func(*SynonymsDeleteSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonymRule) WithFilterPath(v ...string) func(*SynonymsDeleteSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonymRule) WithHeader(h map[string]string) func(*SynonymsDeleteSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SynonymsDeleteSynonymRule) WithOpaqueID(s string) func(*SynonymsDeleteSynonymRuleRequest) {
	_ = "STUB: not implemented"
	return nil
}
