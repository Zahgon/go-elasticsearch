package types

type QueryRulesetMatchedRule struct {
	RuleId string `json:"rule_id"`

	RulesetId string `json:"ruleset_id"`
}

func (s *QueryRulesetMatchedRule) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewQueryRulesetMatchedRule() *QueryRulesetMatchedRule { _ = "STUB: not implemented"; return nil }
