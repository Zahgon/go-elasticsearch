package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _holtWintersMovingAverageAggregation struct {
	v *types.HoltWintersMovingAverageAggregation
}

func NewHoltWintersMovingAverageAggregation(settings types.HoltWintersModelSettingsVariant) *_holtWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersMovingAverageAggregation) Settings(settings types.HoltWintersModelSettingsVariant) *_holtWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersMovingAverageAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_holtWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersMovingAverageAggregation) Format(format string) *_holtWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersMovingAverageAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_holtWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersMovingAverageAggregation) Minimize(minimize bool) *_holtWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersMovingAverageAggregation) Predict(predict int) *_holtWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersMovingAverageAggregation) Window(window int) *_holtWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersMovingAverageAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_holtWintersMovingAverageAggregation) HoltWintersMovingAverageAggregationCaster() *types.HoltWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
