package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/childscoremode"
)

type _nestedQuery struct {
	v *types.NestedQuery
}

func NewNestedQuery(query types.QueryVariant) *_nestedQuery { _ = "STUB: not implemented"; return nil }

func (s *_nestedQuery) IgnoreUnmapped(ignoreunmapped bool) *_nestedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedQuery) InnerHits(innerhits types.InnerHitsVariant) *_nestedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedQuery) Path(field string) *_nestedQuery { _ = "STUB: not implemented"; return nil }

func (s *_nestedQuery) Query(query types.QueryVariant) *_nestedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedQuery) ScoreMode(scoremode childscoremode.ChildScoreMode) *_nestedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedQuery) Boost(boost float32) *_nestedQuery { _ = "STUB: not implemented"; return nil }

func (s *_nestedQuery) QueryName_(queryname_ string) *_nestedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_nestedQuery) NestedQueryCaster() *types.NestedQuery {
	_ = "STUB: not implemented"
	return nil
}
