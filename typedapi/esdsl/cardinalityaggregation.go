package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cardinalityexecutionmode"
)

type _cardinalityAggregation struct {
	v *types.CardinalityAggregation
}

func NewCardinalityAggregation() *_cardinalityAggregation { _ = "STUB: not implemented"; return nil }

func (s *_cardinalityAggregation) ExecutionHint(executionhint cardinalityexecutionmode.CardinalityExecutionMode) *_cardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cardinalityAggregation) PrecisionThreshold(precisionthreshold int) *_cardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cardinalityAggregation) Rehash(rehash bool) *_cardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cardinalityAggregation) Field(field string) *_cardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cardinalityAggregation) Missing(missing types.MissingVariant) *_cardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cardinalityAggregation) Script(script types.ScriptVariant) *_cardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cardinalityAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cardinalityAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cardinalityAggregation) CardinalityAggregationCaster() *types.CardinalityAggregation {
	_ = "STUB: not implemented"
	return nil
}
