package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type CompositeHistogramAggregation struct {
	Field         *string                    `json:"field,omitempty"`
	Interval      Float64                    `json:"interval"`
	MissingBucket *bool                      `json:"missing_bucket,omitempty"`
	MissingOrder  *missingorder.MissingOrder `json:"missing_order,omitempty"`
	Order         *sortorder.SortOrder       `json:"order,omitempty"`

	Script    *Script              `json:"script,omitempty"`
	ValueType *valuetype.ValueType `json:"value_type,omitempty"`
}

func (s *CompositeHistogramAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompositeHistogramAggregation() *CompositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

type CompositeHistogramAggregationVariant interface {
	CompositeHistogramAggregationCaster() *CompositeHistogramAggregation
}

func (s *CompositeHistogramAggregation) CompositeHistogramAggregationCaster() *CompositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
