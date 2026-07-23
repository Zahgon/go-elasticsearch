package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tdigestexecutionhint"
)

type _boxplotAggregation struct {
	v *types.BoxplotAggregation
}

func NewBoxplotAggregation() *_boxplotAggregation { _ = "STUB: not implemented"; return nil }

func (s *_boxplotAggregation) Compression(compression types.Float64) *_boxplotAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boxplotAggregation) ExecutionHint(executionhint tdigestexecutionhint.TDigestExecutionHint) *_boxplotAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boxplotAggregation) Field(field string) *_boxplotAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boxplotAggregation) Missing(missing types.MissingVariant) *_boxplotAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boxplotAggregation) Script(script types.ScriptVariant) *_boxplotAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boxplotAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_boxplotAggregation) BoxplotAggregationCaster() *types.BoxplotAggregation {
	_ = "STUB: not implemented"
	return nil
}
