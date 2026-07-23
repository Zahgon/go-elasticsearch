package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type DateRangeQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Format *string `json:"format,omitempty"`

	Gt *string `json:"gt,omitempty"`

	Gte *string `json:"gte,omitempty"`

	Lt *string `json:"lt,omitempty"`

	Lte        *string `json:"lte,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`

	Relation *rangerelation.RangeRelation `json:"relation,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *DateRangeQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDateRangeQuery() *DateRangeQuery { _ = "STUB: not implemented"; return nil }

type DateRangeQueryVariant interface {
	DateRangeQueryCaster() *DateRangeQuery
}

func (s *DateRangeQuery) DateRangeQueryCaster() *DateRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *DateRangeQuery) RangeQueryCaster() *RangeQuery { _ = "STUB: not implemented"; return nil }
