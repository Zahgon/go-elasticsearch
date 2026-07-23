package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _linearMovingAverageAggregation struct {
	v *types.LinearMovingAverageAggregation
}

func NewLinearMovingAverageAggregation(settings types.EmptyObjectVariant) *_linearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearMovingAverageAggregation) Settings(settings types.EmptyObjectVariant) *_linearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearMovingAverageAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_linearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearMovingAverageAggregation) Format(format string) *_linearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearMovingAverageAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_linearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearMovingAverageAggregation) Minimize(minimize bool) *_linearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearMovingAverageAggregation) Predict(predict int) *_linearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearMovingAverageAggregation) Window(window int) *_linearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearMovingAverageAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearMovingAverageAggregation) LinearMovingAverageAggregationCaster() *types.LinearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
