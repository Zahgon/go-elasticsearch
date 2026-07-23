package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _constantScoreQuery struct {
	v *types.ConstantScoreQuery
}

func NewConstantScoreQuery(filter types.QueryVariant) *_constantScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantScoreQuery) Filter(filter types.QueryVariant) *_constantScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantScoreQuery) Boost(boost float32) *_constantScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantScoreQuery) QueryName_(queryname_ string) *_constantScoreQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_constantScoreQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_constantScoreQuery) ConstantScoreQueryCaster() *types.ConstantScoreQuery {
	_ = "STUB: not implemented"
	return nil
}
