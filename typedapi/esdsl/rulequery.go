package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _ruleQuery struct {
	v *types.RuleQuery
}

func NewRuleQuery(matchcriteria json.RawMessage, organic types.QueryVariant) *_ruleQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleQuery) MatchCriteria(matchcriteria json.RawMessage) *_ruleQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleQuery) Organic(organic types.QueryVariant) *_ruleQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleQuery) RulesetId(rulesetid string) *_ruleQuery { _ = "STUB: not implemented"; return nil }

func (s *_ruleQuery) RulesetIds(rulesetids ...string) *_ruleQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleQuery) Boost(boost float32) *_ruleQuery { _ = "STUB: not implemented"; return nil }

func (s *_ruleQuery) QueryName_(queryname_ string) *_ruleQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_ruleQuery) RuleQueryCaster() *types.RuleQuery { _ = "STUB: not implemented"; return nil }
