package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _ewmaMovingAverageAggregation struct {
	v *types.EwmaMovingAverageAggregation
}

func NewEwmaMovingAverageAggregation(settings types.EwmaModelSettingsVariant) *_ewmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaMovingAverageAggregation) Settings(settings types.EwmaModelSettingsVariant) *_ewmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaMovingAverageAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_ewmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaMovingAverageAggregation) Format(format string) *_ewmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaMovingAverageAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_ewmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaMovingAverageAggregation) Minimize(minimize bool) *_ewmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaMovingAverageAggregation) Predict(predict int) *_ewmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaMovingAverageAggregation) Window(window int) *_ewmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaMovingAverageAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ewmaMovingAverageAggregation) EwmaMovingAverageAggregationCaster() *types.EwmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
