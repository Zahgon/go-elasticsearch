package esapi

import (
	"context"
	"net/http"
)

func newQueryRulesListRulesetsFunc(t Transport) QueryRulesListRulesets {
	_ = "STUB: not implemented"
	return *new(QueryRulesListRulesets)
}

type QueryRulesListRulesets func(o ...func(*QueryRulesListRulesetsRequest)) (*Response, error)

type QueryRulesListRulesetsRequest struct {
	From *int
	Size *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r QueryRulesListRulesetsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f QueryRulesListRulesets) WithContext(v context.Context) func(*QueryRulesListRulesetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesListRulesets) WithFrom(v int) func(*QueryRulesListRulesetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesListRulesets) WithSize(v int) func(*QueryRulesListRulesetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesListRulesets) WithPretty() func(*QueryRulesListRulesetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesListRulesets) WithHuman() func(*QueryRulesListRulesetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesListRulesets) WithErrorTrace() func(*QueryRulesListRulesetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesListRulesets) WithFilterPath(v ...string) func(*QueryRulesListRulesetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesListRulesets) WithHeader(h map[string]string) func(*QueryRulesListRulesetsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesListRulesets) WithOpaqueID(s string) func(*QueryRulesListRulesetsRequest) {
	_ = "STUB: not implemented"
	return nil
}
