package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _minAggregation struct {
	v *types.MinAggregation
}

func NewMinAggregation() *_minAggregation { _ = "STUB: not implemented"; return nil }

func (s *_minAggregation) Field(field string) *_minAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minAggregation) Format(format string) *_minAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minAggregation) Missing(missing types.MissingVariant) *_minAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minAggregation) Script(script types.ScriptVariant) *_minAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minAggregation) MinAggregationCaster() *types.MinAggregation {
	_ = "STUB: not implemented"
	return nil
}
