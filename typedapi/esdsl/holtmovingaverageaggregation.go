package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _holtMovingAverageAggregation struct {
	v *types.HoltMovingAverageAggregation
}

func NewHoltMovingAverageAggregation(settings types.HoltLinearModelSettingsVariant) *_holtMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtMovingAverageAggregation) Settings(settings types.HoltLinearModelSettingsVariant) *_holtMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtMovingAverageAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_holtMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtMovingAverageAggregation) Format(format string) *_holtMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtMovingAverageAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_holtMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtMovingAverageAggregation) Minimize(minimize bool) *_holtMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtMovingAverageAggregation) Predict(predict int) *_holtMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtMovingAverageAggregation) Window(window int) *_holtMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtMovingAverageAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtMovingAverageAggregation) HoltMovingAverageAggregationCaster() *types.HoltMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
