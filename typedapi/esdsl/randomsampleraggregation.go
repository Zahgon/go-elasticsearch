package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _randomSamplerAggregation struct {
	v *types.RandomSamplerAggregation
}

func NewRandomSamplerAggregation(probability types.Float64) *_randomSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_randomSamplerAggregation) Probability(probability types.Float64) *_randomSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_randomSamplerAggregation) Seed(seed int) *_randomSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_randomSamplerAggregation) ShardSeed(shardseed int) *_randomSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_randomSamplerAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_randomSamplerAggregation) RandomSamplerAggregationCaster() *types.RandomSamplerAggregation {
	_ = "STUB: not implemented"
	return nil
}
