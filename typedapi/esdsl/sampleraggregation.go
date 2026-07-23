package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _samplerAggregation struct {
	v *types.SamplerAggregation
}

func NewSamplerAggregation() *_samplerAggregation { _ = "STUB: not implemented"; return nil }

func (s *_samplerAggregation) ShardSize(shardsize int) *_samplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_samplerAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_samplerAggregation) SamplerAggregationCaster() *types.SamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}
