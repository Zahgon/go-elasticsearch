package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sumAggregation struct {
	v *types.SumAggregation
}

func NewSumAggregation() *_sumAggregation { _ = "STUB: not implemented"; return nil }

func (s *_sumAggregation) Field(field string) *_sumAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sumAggregation) Format(format string) *_sumAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sumAggregation) Missing(missing types.MissingVariant) *_sumAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sumAggregation) Script(script types.ScriptVariant) *_sumAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sumAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sumAggregation) SumAggregationCaster() *types.SumAggregation {
	_ = "STUB: not implemented"
	return nil
}
