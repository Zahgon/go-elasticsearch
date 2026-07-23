package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/appliesto"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conditionoperator"
)

type _ruleCondition struct {
	v *types.RuleCondition
}

func NewRuleCondition(appliesto appliesto.AppliesTo, operator conditionoperator.ConditionOperator, value types.Float64) *_ruleCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleCondition) AppliesTo(appliesto appliesto.AppliesTo) *_ruleCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleCondition) Operator(operator conditionoperator.ConditionOperator) *_ruleCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleCondition) Value(value types.Float64) *_ruleCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleCondition) RuleConditionCaster() *types.RuleCondition {
	_ = "STUB: not implemented"
	return nil
}
