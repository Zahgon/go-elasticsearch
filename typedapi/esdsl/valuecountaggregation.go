package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _valueCountAggregation struct {
	v *types.ValueCountAggregation
}

func NewValueCountAggregation() *_valueCountAggregation { _ = "STUB: not implemented"; return nil }

func (s *_valueCountAggregation) Field(field string) *_valueCountAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_valueCountAggregation) Format(format string) *_valueCountAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_valueCountAggregation) Missing(missing types.MissingVariant) *_valueCountAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_valueCountAggregation) Script(script types.ScriptVariant) *_valueCountAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_valueCountAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_valueCountAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_valueCountAggregation) ValueCountAggregationCaster() *types.ValueCountAggregation {
	_ = "STUB: not implemented"
	return nil
}
