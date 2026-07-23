package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _cumulativeCardinalityAggregation struct {
	v *types.CumulativeCardinalityAggregation
}

func NewCumulativeCardinalityAggregation() *_cumulativeCardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeCardinalityAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_cumulativeCardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeCardinalityAggregation) Format(format string) *_cumulativeCardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeCardinalityAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_cumulativeCardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeCardinalityAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeCardinalityAggregation) CumulativeCardinalityAggregationCaster() *types.CumulativeCardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}
