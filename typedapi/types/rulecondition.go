package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/appliesto"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conditionoperator"
)

type RuleCondition struct {
	AppliesTo appliesto.AppliesTo `json:"applies_to"`

	Operator conditionoperator.ConditionOperator `json:"operator"`

	Value Float64 `json:"value"`
}

func (s *RuleCondition) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRuleCondition() *RuleCondition { _ = "STUB: not implemented"; return nil }

type RuleConditionVariant interface {
	RuleConditionCaster() *RuleCondition
}

func (s *RuleCondition) RuleConditionCaster() *RuleCondition { _ = "STUB: not implemented"; return nil }
