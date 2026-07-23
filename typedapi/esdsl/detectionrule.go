package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ruleaction"
)

type _detectionRule struct {
	v *types.DetectionRule
}

func NewDetectionRule() *_detectionRule { _ = "STUB: not implemented"; return nil }

func (s *_detectionRule) Actions(actions ...ruleaction.RuleAction) *_detectionRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detectionRule) Conditions(conditions ...types.RuleConditionVariant) *_detectionRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detectionRule) ConditionsValues(conditionsvalues []types.RuleCondition) *_detectionRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detectionRule) Scope(scope map[string]types.FilterRef) *_detectionRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detectionRule) AddScope(key string, value types.FilterRefVariant) *_detectionRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_detectionRule) DetectionRuleCaster() *types.DetectionRule {
	_ = "STUB: not implemented"
	return nil
}
