package types

type FilteringRules struct {
	AdvancedSnippet FilteringAdvancedSnippet `json:"advanced_snippet"`
	Rules           []FilteringRule          `json:"rules"`
	Validation      FilteringRulesValidation `json:"validation"`
}

func NewFilteringRules() *FilteringRules { _ = "STUB: not implemented"; return nil }

type FilteringRulesVariant interface {
	FilteringRulesCaster() *FilteringRules
}

func (s *FilteringRules) FilteringRulesCaster() *FilteringRules {
	_ = "STUB: not implemented"
	return nil
}
