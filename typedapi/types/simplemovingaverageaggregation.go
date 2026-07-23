package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type SimpleMovingAverageAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Minimize  *bool                `json:"minimize,omitempty"`
	Model     string               `json:"model,omitempty"`
	Predict   *int                 `json:"predict,omitempty"`
	Settings  EmptyObject          `json:"settings"`
	Window    *int                 `json:"window,omitempty"`
}

func (s *SimpleMovingAverageAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SimpleMovingAverageAggregation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSimpleMovingAverageAggregation() *SimpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

type SimpleMovingAverageAggregationVariant interface {
	SimpleMovingAverageAggregationCaster() *SimpleMovingAverageAggregation
}

func (s *SimpleMovingAverageAggregation) SimpleMovingAverageAggregationCaster() *SimpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimpleMovingAverageAggregation) MovingAverageAggregationCaster() *MovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
