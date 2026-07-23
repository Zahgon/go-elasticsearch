package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _filteringRules struct {
	v *types.FilteringRules
}

func NewFilteringRules(advancedsnippet types.FilteringAdvancedSnippetVariant, validation types.FilteringRulesValidationVariant) *_filteringRules {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRules) AdvancedSnippet(advancedsnippet types.FilteringAdvancedSnippetVariant) *_filteringRules {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRules) Rules(rules ...types.FilteringRuleVariant) *_filteringRules {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRules) RulesValues(rulesvalues []types.FilteringRule) *_filteringRules {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRules) Validation(validation types.FilteringRulesValidationVariant) *_filteringRules {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRules) FilteringRulesCaster() *types.FilteringRules {
	_ = "STUB: not implemented"
	return nil
}
