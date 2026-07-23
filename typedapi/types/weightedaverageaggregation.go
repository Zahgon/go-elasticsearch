package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type WeightedAverageAggregation struct {
	Format *string `json:"format,omitempty"`

	Value     *WeightedAverageValue `json:"value,omitempty"`
	ValueType *valuetype.ValueType  `json:"value_type,omitempty"`

	Weight *WeightedAverageValue `json:"weight,omitempty"`
}

func (s *WeightedAverageAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewWeightedAverageAggregation() *WeightedAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

type WeightedAverageAggregationVariant interface {
	WeightedAverageAggregationCaster() *WeightedAverageAggregation
}

func (s *WeightedAverageAggregation) WeightedAverageAggregationCaster() *WeightedAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
