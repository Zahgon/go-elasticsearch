package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type _dateRangeQuery struct {
	k string
	v *types.DateRangeQuery
}

func NewDateRangeQuery(key string) *_dateRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_dateRangeQuery) Format(dateformat string) *_dateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) TimeZone(timezone string) *_dateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) Boost(boost float32) *_dateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) Gt(datemath string) *_dateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) Gte(datemath string) *_dateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) Lt(datemath string) *_dateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) Lte(datemath string) *_dateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) QueryName_(queryname_ string) *_dateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) Relation(relation rangerelation.RangeRelation) *_dateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_dateRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewSingleDateRangeQuery() *_dateRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_dateRangeQuery) DateRangeQueryCaster() *types.DateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}
