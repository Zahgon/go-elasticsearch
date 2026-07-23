package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type _longNumberRangeQuery struct {
	k string
	v *types.LongNumberRangeQuery
}

func NewLongNumberRangeQuery(key string) *_longNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) Boost(boost float32) *_longNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) Gt(gt int64) *_longNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) Gte(gte int64) *_longNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) Lt(lt int64) *_longNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) Lte(lte int64) *_longNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) QueryName_(queryname_ string) *_longNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) Relation(relation rangerelation.RangeRelation) *_longNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_longNumberRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_longNumberRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewSingleLongNumberRangeQuery() *_longNumberRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_longNumberRangeQuery) LongNumberRangeQueryCaster() *types.LongNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}
