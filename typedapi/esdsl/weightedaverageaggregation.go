package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type _weightedAverageAggregation struct {
	v *types.WeightedAverageAggregation
}

func NewWeightedAverageAggregation() *_weightedAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedAverageAggregation) Format(format string) *_weightedAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedAverageAggregation) Value(value types.WeightedAverageValueVariant) *_weightedAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedAverageAggregation) ValueType(valuetype valuetype.ValueType) *_weightedAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedAverageAggregation) Weight(weight types.WeightedAverageValueVariant) *_weightedAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedAverageAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedAverageAggregation) WeightedAverageAggregationCaster() *types.WeightedAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
