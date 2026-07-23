package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ruleaction"
)

type DetectionRule struct {
	Actions []ruleaction.RuleAction `json:"actions,omitempty"`

	Conditions []RuleCondition `json:"conditions,omitempty"`

	Scope map[string]FilterRef `json:"scope,omitempty"`
}

func NewDetectionRule() *DetectionRule { _ = "STUB: not implemented"; return nil }

type DetectionRuleVariant interface {
	DetectionRuleCaster() *DetectionRule
}

func (s *DetectionRule) DetectionRuleCaster() *DetectionRule { _ = "STUB: not implemented"; return nil }
