package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hasParentQuery struct {
	v *types.HasParentQuery
}

func NewHasParentQuery(query types.QueryVariant) *_hasParentQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasParentQuery) IgnoreUnmapped(ignoreunmapped bool) *_hasParentQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasParentQuery) InnerHits(innerhits types.InnerHitsVariant) *_hasParentQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasParentQuery) ParentType(relationname string) *_hasParentQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasParentQuery) Query(query types.QueryVariant) *_hasParentQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasParentQuery) Score(score bool) *_hasParentQuery { _ = "STUB: not implemented"; return nil }

func (s *_hasParentQuery) Boost(boost float32) *_hasParentQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasParentQuery) QueryName_(queryname_ string) *_hasParentQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hasParentQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_hasParentQuery) HasParentQueryCaster() *types.HasParentQuery {
	_ = "STUB: not implemented"
	return nil
}
