package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type LinearMovingAverageAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Minimize  *bool                `json:"minimize,omitempty"`
	Model     string               `json:"model,omitempty"`
	Predict   *int                 `json:"predict,omitempty"`
	Settings  EmptyObject          `json:"settings"`
	Window    *int                 `json:"window,omitempty"`
}

func (s *LinearMovingAverageAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s LinearMovingAverageAggregation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLinearMovingAverageAggregation() *LinearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

type LinearMovingAverageAggregationVariant interface {
	LinearMovingAverageAggregationCaster() *LinearMovingAverageAggregation
}

func (s *LinearMovingAverageAggregation) LinearMovingAverageAggregationCaster() *LinearMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *LinearMovingAverageAggregation) MovingAverageAggregationCaster() *MovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
