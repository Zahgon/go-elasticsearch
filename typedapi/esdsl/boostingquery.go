package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _boostingQuery struct {
	v *types.BoostingQuery
}

func NewBoostingQuery(negative types.QueryVariant, negativeboost types.Float64, positive types.QueryVariant) *_boostingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boostingQuery) Negative(negative types.QueryVariant) *_boostingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boostingQuery) NegativeBoost(negativeboost types.Float64) *_boostingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boostingQuery) Positive(positive types.QueryVariant) *_boostingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boostingQuery) Boost(boost float32) *_boostingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boostingQuery) QueryName_(queryname_ string) *_boostingQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boostingQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_boostingQuery) BoostingQueryCaster() *types.BoostingQuery {
	_ = "STUB: not implemented"
	return nil
}
