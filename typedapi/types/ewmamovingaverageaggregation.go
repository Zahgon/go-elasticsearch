package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type EwmaMovingAverageAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Minimize  *bool                `json:"minimize,omitempty"`
	Model     string               `json:"model,omitempty"`
	Predict   *int                 `json:"predict,omitempty"`
	Settings  EwmaModelSettings    `json:"settings"`
	Window    *int                 `json:"window,omitempty"`
}

func (s *EwmaMovingAverageAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s EwmaMovingAverageAggregation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewEwmaMovingAverageAggregation() *EwmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

type EwmaMovingAverageAggregationVariant interface {
	EwmaMovingAverageAggregationCaster() *EwmaMovingAverageAggregation
}

func (s *EwmaMovingAverageAggregation) EwmaMovingAverageAggregationCaster() *EwmaMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *EwmaMovingAverageAggregation) MovingAverageAggregationCaster() *MovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
