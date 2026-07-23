package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type NumberRangeQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Gt *Float64 `json:"gt,omitempty"`

	Gte *Float64 `json:"gte,omitempty"`

	Lt *Float64 `json:"lt,omitempty"`

	Lte        *Float64 `json:"lte,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`

	Relation *rangerelation.RangeRelation `json:"relation,omitempty"`
}

func (s *NumberRangeQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNumberRangeQuery() *NumberRangeQuery { _ = "STUB: not implemented"; return nil }

type NumberRangeQueryVariant interface {
	NumberRangeQueryCaster() *NumberRangeQuery
}

func (s *NumberRangeQuery) NumberRangeQueryCaster() *NumberRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *NumberRangeQuery) RangeQueryCaster() *RangeQuery { _ = "STUB: not implemented"; return nil }
