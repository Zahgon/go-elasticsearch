package types

type QueryRulesetListItem struct {
	RuleCriteriaTypesCounts map[string]int `json:"rule_criteria_types_counts"`

	RuleTotalCount int `json:"rule_total_count"`

	RuleTypeCounts map[string]int `json:"rule_type_counts"`

	RulesetId string `json:"ruleset_id"`
}

func (s *QueryRulesetListItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewQueryRulesetListItem() *QueryRulesetListItem { _ = "STUB: not implemented"; return nil }
