package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/childscoremode"
)

type _hasChildQuery struct {
	v *types.HasChildQuery
}

func NewHasChildQuery(query types.QueryVariant) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) IgnoreUnmapped(ignoreunmapped bool) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) InnerHits(innerhits types.InnerHitsVariant) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) MaxChildren(maxchildren int) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) MinChildren(minchildren int) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) Query(query types.QueryVariant) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) ScoreMode(scoremode childscoremode.ChildScoreMode) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) Type(relationname string) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) Boost(boost float32) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) QueryName_(queryname_ string) *_hasChildQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasChildQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_hasChildQuery) HasChildQueryCaster() *types.HasChildQuery {
	_ = "STUB: not implemented"
	return nil
}
