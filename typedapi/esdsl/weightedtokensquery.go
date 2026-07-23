package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _weightedTokensQuery struct {
	k string
	v *types.WeightedTokensQuery
}

func NewWeightedTokensQuery(key string) *_weightedTokensQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedTokensQuery) PruningConfig(pruningconfig types.TokenPruningConfigVariant) *_weightedTokensQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedTokensQuery) Tokens(tokens []map[string]float32) *_weightedTokensQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedTokensQuery) Boost(boost float32) *_weightedTokensQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedTokensQuery) QueryName_(queryname_ string) *_weightedTokensQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedTokensQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func NewSingleWeightedTokensQuery() *_weightedTokensQuery { _ = "STUB: not implemented"; return nil }

func (s *_weightedTokensQuery) WeightedTokensQueryCaster() *types.WeightedTokensQuery {
	_ = "STUB: not implemented"
	return nil
}
