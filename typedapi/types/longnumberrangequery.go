package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type LongNumberRangeQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Gt *int64 `json:"gt,omitempty"`

	Gte *int64 `json:"gte,omitempty"`

	Lt *int64 `json:"lt,omitempty"`

	Lte        *int64  `json:"lte,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`

	Relation *rangerelation.RangeRelation `json:"relation,omitempty"`
}

func (s *LongNumberRangeQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLongNumberRangeQuery() *LongNumberRangeQuery { _ = "STUB: not implemented"; return nil }

type LongNumberRangeQueryVariant interface {
	LongNumberRangeQueryCaster() *LongNumberRangeQuery
}

func (s *LongNumberRangeQuery) LongNumberRangeQueryCaster() *LongNumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *LongNumberRangeQuery) RangeQueryCaster() *RangeQuery {
	_ = "STUB: not implemented"
	return nil
}
