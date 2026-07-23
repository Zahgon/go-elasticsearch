package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type TermRangeQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Gt *string `json:"gt,omitempty"`

	Gte *string `json:"gte,omitempty"`

	Lt *string `json:"lt,omitempty"`

	Lte        *string `json:"lte,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`

	Relation *rangerelation.RangeRelation `json:"relation,omitempty"`
}

func (s *TermRangeQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermRangeQuery() *TermRangeQuery { _ = "STUB: not implemented"; return nil }

type TermRangeQueryVariant interface {
	TermRangeQueryCaster() *TermRangeQuery
}

func (s *TermRangeQuery) TermRangeQueryCaster() *TermRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *TermRangeQuery) RangeQueryCaster() *RangeQuery { _ = "STUB: not implemented"; return nil }
