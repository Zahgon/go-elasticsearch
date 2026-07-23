package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _timeSeriesAggregation struct {
	v *types.TimeSeriesAggregation
}

func NewTimeSeriesAggregation() *_timeSeriesAggregation { _ = "STUB: not implemented"; return nil }

func (s *_timeSeriesAggregation) Keyed(keyed bool) *_timeSeriesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_timeSeriesAggregation) Size(size int) *_timeSeriesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_timeSeriesAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_timeSeriesAggregation) TimeSeriesAggregationCaster() *types.TimeSeriesAggregation {
	_ = "STUB: not implemented"
	return nil
}
