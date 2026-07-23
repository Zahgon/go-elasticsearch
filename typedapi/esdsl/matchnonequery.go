package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _matchNoneQuery struct {
	v *types.MatchNoneQuery
}

func NewMatchNoneQuery() *_matchNoneQuery { _ = "STUB: not implemented"; return nil }

func (s *_matchNoneQuery) Boost(boost float32) *_matchNoneQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchNoneQuery) QueryName_(queryname_ string) *_matchNoneQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_matchNoneQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_matchNoneQuery) MatchNoneQueryCaster() *types.MatchNoneQuery {
	_ = "STUB: not implemented"
	return nil
}
