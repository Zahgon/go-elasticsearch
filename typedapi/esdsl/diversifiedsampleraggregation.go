package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sampleraggregationexecutionhint"
)

type _diversifiedSamplerAggregation struct {
	v *types.DiversifiedSamplerAggregation
}

func NewDiversifiedSamplerAggregation() *_diversifiedSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifiedSamplerAggregation) ExecutionHint(executionhint sampleraggregationexecutionhint.SamplerAggregationExecutionHint) *_diversifiedSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifiedSamplerAggregation) Field(field string) *_diversifiedSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifiedSamplerAggregation) MaxDocsPerValue(maxdocspervalue int) *_diversifiedSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifiedSamplerAggregation) Script(script types.ScriptVariant) *_diversifiedSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifiedSamplerAggregation) ShardSize(shardsize int) *_diversifiedSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifiedSamplerAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifiedSamplerAggregation) DiversifiedSamplerAggregationCaster() *types.DiversifiedSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}
