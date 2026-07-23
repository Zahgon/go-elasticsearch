package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tdigestexecutionhint"
)

type _medianAbsoluteDeviationAggregation struct {
	v *types.MedianAbsoluteDeviationAggregation
}

func NewMedianAbsoluteDeviationAggregation() *_medianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_medianAbsoluteDeviationAggregation) Compression(compression types.Float64) *_medianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_medianAbsoluteDeviationAggregation) ExecutionHint(executionhint tdigestexecutionhint.TDigestExecutionHint) *_medianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_medianAbsoluteDeviationAggregation) Field(field string) *_medianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_medianAbsoluteDeviationAggregation) Format(format string) *_medianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_medianAbsoluteDeviationAggregation) Missing(missing types.MissingVariant) *_medianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_medianAbsoluteDeviationAggregation) Script(script types.ScriptVariant) *_medianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_medianAbsoluteDeviationAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_medianAbsoluteDeviationAggregation) MedianAbsoluteDeviationAggregationCaster() *types.MedianAbsoluteDeviationAggregation {
	_ = "STUB: not implemented"
	return nil
}
