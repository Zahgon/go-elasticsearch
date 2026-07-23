package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type _termRangeQuery struct {
	k string
	v *types.TermRangeQuery
}

func NewTermRangeQuery(key string) *_termRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_termRangeQuery) Boost(boost float32) *_termRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termRangeQuery) Gt(gt string) *_termRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_termRangeQuery) Gte(gte string) *_termRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_termRangeQuery) Lt(lt string) *_termRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_termRangeQuery) Lte(lte string) *_termRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_termRangeQuery) QueryName_(queryname_ string) *_termRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termRangeQuery) Relation(relation rangerelation.RangeRelation) *_termRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termRangeQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_termRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewSingleTermRangeQuery() *_termRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_termRangeQuery) TermRangeQueryCaster() *types.TermRangeQuery {
	_ = "STUB: not implemented"
	return nil
}
