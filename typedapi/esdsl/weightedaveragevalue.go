package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _weightedAverageValue struct {
	v *types.WeightedAverageValue
}

func NewWeightedAverageValue() *_weightedAverageValue { _ = "STUB: not implemented"; return nil }

func (s *_weightedAverageValue) Field(field string) *_weightedAverageValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedAverageValue) Missing(missing types.Float64) *_weightedAverageValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedAverageValue) Script(script types.ScriptVariant) *_weightedAverageValue {
	_ = "STUB: not implemented"
	return nil
}

func (s *_weightedAverageValue) WeightedAverageValueCaster() *types.WeightedAverageValue {
	_ = "STUB: not implemented"
	return nil
}
