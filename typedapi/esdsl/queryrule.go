package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/queryruletype"
)

type _queryRule struct {
	v *types.QueryRule
}

func NewQueryRule(actions types.QueryRuleActionsVariant, type_ queryruletype.QueryRuleType) *_queryRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRule) Actions(actions types.QueryRuleActionsVariant) *_queryRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRule) Criteria(criteria ...types.QueryRuleCriteriaVariant) *_queryRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRule) Priority(priority int) *_queryRule { _ = "STUB: not implemented"; return nil }

func (s *_queryRule) RuleId(id string) *_queryRule { _ = "STUB: not implemented"; return nil }

func (s *_queryRule) Type(type_ queryruletype.QueryRuleType) *_queryRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryRule) QueryRuleCaster() *types.QueryRule { _ = "STUB: not implemented"; return nil }
