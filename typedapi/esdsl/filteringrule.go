package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filteringpolicy"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/filteringrulerule"
)

type _filteringRule struct {
	v *types.FilteringRule
}

func NewFilteringRule(order int, policy filteringpolicy.FilteringPolicy, rule filteringrulerule.FilteringRuleRule, value string) *_filteringRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRule) CreatedAt(datetime types.DateTimeVariant) *_filteringRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRule) Field(field string) *_filteringRule { _ = "STUB: not implemented"; return nil }

func (s *_filteringRule) Id(id string) *_filteringRule { _ = "STUB: not implemented"; return nil }

func (s *_filteringRule) Order(order int) *_filteringRule { _ = "STUB: not implemented"; return nil }

func (s *_filteringRule) Policy(policy filteringpolicy.FilteringPolicy) *_filteringRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRule) Rule(rule filteringrulerule.FilteringRuleRule) *_filteringRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRule) UpdatedAt(datetime types.DateTimeVariant) *_filteringRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringRule) Value(value string) *_filteringRule { _ = "STUB: not implemented"; return nil }

func (s *_filteringRule) FilteringRuleCaster() *types.FilteringRule {
	_ = "STUB: not implemented"
	return nil
}
