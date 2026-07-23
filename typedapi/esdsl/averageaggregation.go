package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _averageAggregation struct {
	v *types.AverageAggregation
}

func NewAverageAggregation() *_averageAggregation { _ = "STUB: not implemented"; return nil }

func (s *_averageAggregation) Field(field string) *_averageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageAggregation) Format(format string) *_averageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageAggregation) Missing(missing types.MissingVariant) *_averageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageAggregation) Script(script types.ScriptVariant) *_averageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageAggregation) AverageAggregationCaster() *types.AverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
