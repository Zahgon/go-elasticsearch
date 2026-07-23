package esapi

import (
	"context"
	"io"
	"net/http"
)

func newQueryRulesTestFunc(t Transport) QueryRulesTest {
	_ = "STUB: not implemented"
	return *new(QueryRulesTest)
}

type QueryRulesTest func(body io.Reader, ruleset_id string, o ...func(*QueryRulesTestRequest)) (*Response, error)

type QueryRulesTestRequest struct {
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

func (r QueryRulesTestRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f QueryRulesTest) WithContext(v context.Context) func(*QueryRulesTestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesTest) WithPretty() func(*QueryRulesTestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesTest) WithHuman() func(*QueryRulesTestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesTest) WithErrorTrace() func(*QueryRulesTestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesTest) WithFilterPath(v ...string) func(*QueryRulesTestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesTest) WithHeader(h map[string]string) func(*QueryRulesTestRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f QueryRulesTest) WithOpaqueID(s string) func(*QueryRulesTestRequest) {
	_ = "STUB: not implemented"
	return nil
}
