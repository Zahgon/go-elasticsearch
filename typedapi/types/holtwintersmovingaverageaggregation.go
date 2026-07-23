package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type HoltWintersMovingAverageAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Format *string `json:"format,omitempty"`

	GapPolicy *gappolicy.GapPolicy     `json:"gap_policy,omitempty"`
	Minimize  *bool                    `json:"minimize,omitempty"`
	Model     string                   `json:"model,omitempty"`
	Predict   *int                     `json:"predict,omitempty"`
	Settings  HoltWintersModelSettings `json:"settings"`
	Window    *int                     `json:"window,omitempty"`
}

func (s *HoltWintersMovingAverageAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s HoltWintersMovingAverageAggregation) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHoltWintersMovingAverageAggregation() *HoltWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

type HoltWintersMovingAverageAggregationVariant interface {
	HoltWintersMovingAverageAggregationCaster() *HoltWintersMovingAverageAggregation
}

func (s *HoltWintersMovingAverageAggregation) HoltWintersMovingAverageAggregationCaster() *HoltWintersMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *HoltWintersMovingAverageAggregation) MovingAverageAggregationCaster() *MovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
