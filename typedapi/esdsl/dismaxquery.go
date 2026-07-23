package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _disMaxQuery struct {
	v *types.DisMaxQuery
}

func NewDisMaxQuery() *_disMaxQuery { _ = "STUB: not implemented"; return nil }

func (s *_disMaxQuery) Queries(queries ...types.QueryVariant) *_disMaxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_disMaxQuery) QueriesValues(queriesvalues []types.Query) *_disMaxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_disMaxQuery) TieBreaker(tiebreaker types.Float64) *_disMaxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_disMaxQuery) Boost(boost float32) *_disMaxQuery { _ = "STUB: not implemented"; return nil }

func (s *_disMaxQuery) QueryName_(queryname_ string) *_disMaxQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_disMaxQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_disMaxQuery) DisMaxQueryCaster() *types.DisMaxQuery {
	_ = "STUB: not implemented"
	return nil
}
