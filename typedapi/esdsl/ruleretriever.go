package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _ruleRetriever struct {
	v *types.RuleRetriever
}

func NewRuleRetriever(matchcriteria json.RawMessage, retriever types.RetrieverContainerVariant) *_ruleRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleRetriever) MatchCriteria(matchcriteria json.RawMessage) *_ruleRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleRetriever) RankWindowSize(rankwindowsize int) *_ruleRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleRetriever) Retriever(retriever types.RetrieverContainerVariant) *_ruleRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleRetriever) RulesetIds(rulesetids ...string) *_ruleRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleRetriever) Filter(filters ...types.QueryVariant) *_ruleRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleRetriever) MinScore(minscore float32) *_ruleRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleRetriever) Name_(name_ string) *_ruleRetriever { _ = "STUB: not implemented"; return nil }

func (s *_ruleRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ruleRetriever) RuleRetrieverCaster() *types.RuleRetriever {
	_ = "STUB: not implemented"
	return nil
}
