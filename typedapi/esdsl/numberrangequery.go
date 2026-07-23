package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type _numberRangeQuery struct {
	k string
	v *types.NumberRangeQuery
}

func NewNumberRangeQuery(key string) *_numberRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_numberRangeQuery) Boost(boost float32) *_numberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numberRangeQuery) Gt(gt types.Float64) *_numberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numberRangeQuery) Gte(gte types.Float64) *_numberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numberRangeQuery) Lt(lt types.Float64) *_numberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numberRangeQuery) Lte(lte types.Float64) *_numberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numberRangeQuery) QueryName_(queryname_ string) *_numberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numberRangeQuery) Relation(relation rangerelation.RangeRelation) *_numberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numberRangeQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_numberRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numberRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_numberRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewSingleNumberRangeQuery() *_numberRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_numberRangeQuery) NumberRangeQueryCaster() *types.NumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}
