package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _derivativeAggregation struct {
	v *types.DerivativeAggregation
}

func NewDerivativeAggregation() *_derivativeAggregation { _ = "STUB: not implemented"; return nil }

func (s *_derivativeAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_derivativeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_derivativeAggregation) Format(format string) *_derivativeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_derivativeAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_derivativeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_derivativeAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_derivativeAggregation) DerivativeAggregationCaster() *types.DerivativeAggregation {
	_ = "STUB: not implemented"
	return nil
}
