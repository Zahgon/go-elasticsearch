package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _inferenceAggregation struct {
	v *types.InferenceAggregation
}

func NewInferenceAggregation() *_inferenceAggregation { _ = "STUB: not implemented"; return nil }

func (s *_inferenceAggregation) InferenceConfig(inferenceconfig types.InferenceConfigContainerVariant) *_inferenceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceAggregation) ModelId(name string) *_inferenceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_inferenceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceAggregation) Format(format string) *_inferenceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_inferenceAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceAggregation) InferenceAggregationCaster() *types.InferenceAggregation {
	_ = "STUB: not implemented"
	return nil
}
