package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type CompositeDateHistogramAggregation struct {
	CalendarInterval *string `json:"calendar_interval,omitempty"`

	Field *string `json:"field,omitempty"`

	FixedInterval *string                    `json:"fixed_interval,omitempty"`
	Format        *string                    `json:"format,omitempty"`
	MissingBucket *bool                      `json:"missing_bucket,omitempty"`
	MissingOrder  *missingorder.MissingOrder `json:"missing_order,omitempty"`
	Offset        Duration                   `json:"offset,omitempty"`
	Order         *sortorder.SortOrder       `json:"order,omitempty"`

	Script    *Script              `json:"script,omitempty"`
	TimeZone  *string              `json:"time_zone,omitempty"`
	ValueType *valuetype.ValueType `json:"value_type,omitempty"`
}

func (s *CompositeDateHistogramAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompositeDateHistogramAggregation() *CompositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

type CompositeDateHistogramAggregationVariant interface {
	CompositeDateHistogramAggregationCaster() *CompositeDateHistogramAggregation
}

func (s *CompositeDateHistogramAggregation) CompositeDateHistogramAggregationCaster() *CompositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
