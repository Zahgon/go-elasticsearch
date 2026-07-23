package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _histogramAggregation struct {
	v *types.HistogramAggregation
}

func NewHistogramAggregation() *_histogramAggregation { _ = "STUB: not implemented"; return nil }

func (s *_histogramAggregation) ExtendedBounds(extendedbounds types.ExtendedBoundsdoubleVariant) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) Field(field string) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) Format(format string) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) HardBounds(hardbounds types.ExtendedBoundsdoubleVariant) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) Interval(interval types.Float64) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) Keyed(keyed bool) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) MinDocCount(mindoccount int) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) Missing(missing types.Float64) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) Offset(offset types.Float64) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) Order(aggregateorder types.AggregateOrderVariant) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) Script(script types.ScriptVariant) *_histogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramAggregation) HistogramAggregationCaster() *types.HistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
