package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _synonymRule struct {
	v *types.SynonymRule
}

func NewSynonymRule() *_synonymRule { _ = "STUB: not implemented"; return nil }

func (s *_synonymRule) Id(id string) *_synonymRule { _ = "STUB: not implemented"; return nil }

func (s *_synonymRule) Synonyms(synonymstring string) *_synonymRule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_synonymRule) SynonymRuleCaster() *types.SynonymRule {
	_ = "STUB: not implemented"
	return nil
}
