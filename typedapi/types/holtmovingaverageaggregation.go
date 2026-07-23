package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type HoltMovingAverageAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy    `json:"gap_policy,omitempty"`
	Minimize  *bool                   `json:"minimize,omitempty"`
	Model     string                  `json:"model,omitempty"`
	Predict   *int                    `json:"predict,omitempty"`
	Settings  HoltLinearModelSettings `json:"settings"`
	Window    *int                    `json:"window,omitempty"`
}

func (s *HoltMovingAverageAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s HoltMovingAverageAggregation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHoltMovingAverageAggregation() *HoltMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

type HoltMovingAverageAggregationVariant interface {
	HoltMovingAverageAggregationCaster() *HoltMovingAverageAggregation
}

func (s *HoltMovingAverageAggregation) HoltMovingAverageAggregationCaster() *HoltMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *HoltMovingAverageAggregation) MovingAverageAggregationCaster() *MovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
