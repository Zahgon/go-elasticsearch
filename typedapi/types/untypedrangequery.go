package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type UntypedRangeQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	Format *string `json:"format,omitempty"`

	Gt json.RawMessage `json:"gt,omitempty"`

	Gte json.RawMessage `json:"gte,omitempty"`

	Lt json.RawMessage `json:"lt,omitempty"`

	Lte        json.RawMessage `json:"lte,omitempty"`
	QueryName_ *string         `json:"_name,omitempty"`

	Relation *rangerelation.RangeRelation `json:"relation,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *UntypedRangeQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUntypedRangeQuery() *UntypedRangeQuery { _ = "STUB: not implemented"; return nil }

type UntypedRangeQueryVariant interface {
	UntypedRangeQueryCaster() *UntypedRangeQuery
}

func (s *UntypedRangeQuery) UntypedRangeQueryCaster() *UntypedRangeQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *UntypedRangeQuery) RangeQueryCaster() *RangeQuery { _ = "STUB: not implemented"; return nil }
