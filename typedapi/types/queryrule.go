package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/queryruletype"
)

type QueryRule struct {
	Actions QueryRuleActions `json:"actions"`

	Criteria []QueryRuleCriteria `json:"criteria"`
	Priority *int                `json:"priority,omitempty"`

	RuleId string `json:"rule_id"`

	Type queryruletype.QueryRuleType `json:"type"`
}

func (s *QueryRule) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewQueryRule() *QueryRule { _ = "STUB: not implemented"; return nil }

type QueryRuleVariant interface {
	QueryRuleCaster() *QueryRule
}

func (s *QueryRule) QueryRuleCaster() *QueryRule { _ = "STUB: not implemented"; return nil }
