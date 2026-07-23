package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type _untypedRangeQuery struct {
	k string
	v *types.UntypedRangeQuery
}

func NewUntypedRangeQuery(key string) *_untypedRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_untypedRangeQuery) Format(dateformat string) *_untypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) TimeZone(timezone string) *_untypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) Boost(boost float32) *_untypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) Gt(gt json.RawMessage) *_untypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) Gte(gte json.RawMessage) *_untypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) Lt(lt json.RawMessage) *_untypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) Lte(lte json.RawMessage) *_untypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) QueryName_(queryname_ string) *_untypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) Relation(relation rangerelation.RangeRelation) *_untypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_untypedRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_untypedRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewSingleUntypedRangeQuery() *_untypedRangeQuery { _ = "STUB: not implemented"; return nil }

func (s *_untypedRangeQuery) UntypedRangeQueryCaster() *types.UntypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}
