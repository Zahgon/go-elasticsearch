package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _maxAggregation struct {
	v *types.MaxAggregation
}

func NewMaxAggregation() *_maxAggregation { _ = "STUB: not implemented"; return nil }

func (s *_maxAggregation) Field(field string) *_maxAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxAggregation) Format(format string) *_maxAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxAggregation) Missing(missing types.MissingVariant) *_maxAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxAggregation) Script(script types.ScriptVariant) *_maxAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxAggregation) MaxAggregationCaster() *types.MaxAggregation {
	_ = "STUB: not implemented"
	return nil
}
