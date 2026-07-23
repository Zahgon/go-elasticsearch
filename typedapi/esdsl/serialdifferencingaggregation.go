package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _serialDifferencingAggregation struct {
	v *types.SerialDifferencingAggregation
}

func NewSerialDifferencingAggregation() *_serialDifferencingAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_serialDifferencingAggregation) Lag(lag int) *_serialDifferencingAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_serialDifferencingAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_serialDifferencingAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_serialDifferencingAggregation) Format(format string) *_serialDifferencingAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_serialDifferencingAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_serialDifferencingAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_serialDifferencingAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_serialDifferencingAggregation) SerialDifferencingAggregationCaster() *types.SerialDifferencingAggregation {
	_ = "STUB: not implemented"
	return nil
}
