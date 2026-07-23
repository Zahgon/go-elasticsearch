package types

import (
	"encoding/json"
)

type RuleQuery struct {
	Boost         *float32        `json:"boost,omitempty"`
	MatchCriteria json.RawMessage `json:"match_criteria,omitempty"`
	Organic       Query           `json:"organic"`
	QueryName_    *string         `json:"_name,omitempty"`
	RulesetId     *string         `json:"ruleset_id,omitempty"`
	RulesetIds    []string        `json:"ruleset_ids,omitempty"`
}

func (s *RuleQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRuleQuery() *RuleQuery { _ = "STUB: not implemented"; return nil }

type RuleQueryVariant interface {
	RuleQueryCaster() *RuleQuery
}

func (s *RuleQuery) RuleQueryCaster() *RuleQuery { _ = "STUB: not implemented"; return nil }
