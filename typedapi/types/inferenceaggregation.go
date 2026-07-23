package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type InferenceAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`

	InferenceConfig *InferenceConfigContainer `json:"inference_config,omitempty"`

	ModelId string `json:"model_id"`
}

func (s *InferenceAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceAggregation() *InferenceAggregation { _ = "STUB: not implemented"; return nil }

type InferenceAggregationVariant interface {
	InferenceAggregationCaster() *InferenceAggregation
}

func (s *InferenceAggregation) InferenceAggregationCaster() *InferenceAggregation {
	_ = "STUB: not implemented"
	return nil
}
